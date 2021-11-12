// Package linkie implements the protocol for TP-Link smart home devices with
// xor-encoded communication over port 9999.
//
// TODO: discovery
package linkie

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/pgaskin/kasa/transport"
)

const (
	DevicePort = 9999
	CryptoIV   = byte(0xAB)
)

var (
	DefaultConnectionTimeout = time.Second * 5
	DefaultRequestTimeout    = time.Second * 5
	DefaultResponseTimeout   = time.Second * 15
)

type Client struct {
	IP                net.IP
	ConnectionTimeout time.Duration // timeout for connecting
	RequestTimeout    time.Duration // timeout for writing the request
	ResponseTimeout   time.Duration // timeout for reading the response
}

var _ transport.Requester = (*Client)(nil)

func (c *Client) Request(ctx context.Context, in, out interface{}) error {
	var req encReader
	if b, err := json.Marshal(in); err != nil {
		return fmt.Errorf("encode request: %w", err)
	} else {
		req.Bytes = b
		req.Key = CryptoIV
		req.IncludeSize = out != nil || len(b) > 1024 // i.e. whether to use UDP
	}

	var d net.Dialer
	if c.ConnectionTimeout > 0 {
		d.Timeout = c.ConnectionTimeout
	} else {
		d.Timeout = DefaultConnectionTimeout
	}

	var conn net.Conn
	if req.IncludeSize {
		if x, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: c.IP, Port: 9999}); err != nil {
			return fmt.Errorf("connect (tcp): %w", err)
		} else {
			conn = x
		}
	} else {
		if x, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: c.IP, Port: 9999}); err != nil {
			return fmt.Errorf("connect (udp): %w", err)
		} else {
			conn = x
		}
	}
	defer conn.Close()

	if c.RequestTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(c.RequestTimeout))
	} else {
		conn.SetWriteDeadline(time.Now().Add(DefaultRequestTimeout))
	}

	if _, err := io.Copy(conn, &req); err != nil {
		return fmt.Errorf("write request: %w", err)
	}

	if out != nil {
		var buf bytes.Buffer

		if c.ResponseTimeout > 0 {
			conn.SetReadDeadline(time.Now().Add(c.ResponseTimeout))
		} else {
			conn.SetReadDeadline(time.Now().Add(DefaultResponseTimeout))
		}

		if _, err := buf.ReadFrom(&decReader{
			Source: conn,
			Key:    CryptoIV,
		}); err != nil {
			return fmt.Errorf("read response: %w", err)
		}

		if err := json.Unmarshal(buf.Bytes(), out); err != nil {
			return fmt.Errorf("parse response: %w", err)
		}
	}

	return nil
}

const encSzN = 4

type encReader struct {
	Bytes       []byte
	Key         byte
	IncludeSize bool
	off         int
}

func (r *encReader) Read(b []byte) (n int, err error) {
	t := len(r.Bytes)
	for n < len(b) {
		if r.IncludeSize {
			switch r.off {
			// big-endian length
			case 0:
				b[n] = byte(uint32(t) >> 24)
			case 1:
				b[n] = byte(uint32(t) >> 16)
			case 2:
				b[n] = byte(uint32(t) >> 8)
			case 3:
				b[n] = byte(uint32(t) >> 0)
			default:
				if r.off-4 >= t {
					return n, io.EOF
				}
				b[n] = r.Bytes[r.off-4] ^ r.Key
				r.Key = b[n]
			}
		} else {
			if r.off >= t {
				return n, io.EOF
			}
			b[n] = r.Bytes[r.off] ^ r.Key
			r.Key = b[n]
		}
		r.off++
		n++
	}
	return n, nil
}

type decReader struct {
	Source io.Reader
	Key    byte
	sz     int
	n      int
}

func (d *decReader) Read(p []byte) (n int, err error) {
	if d.n == 0 {
		var x [encSzN]byte
		if _, err := io.ReadFull(d.Source, x[:]); err != nil {
			if err == io.EOF {
				return 0, fmt.Errorf("read size: unexpected eof")
			}
			return 0, fmt.Errorf("read size: %w", err)
		}
		d.sz = int(uint32(x[3]) | uint32(x[2])<<8 | uint32(x[1])<<16 | uint32(x[0])<<24)
		d.n = encSzN
	} else if d.n-encSzN >= d.sz {
		return 0, io.EOF
	}

	if r := d.sz - (d.n - encSzN); len(p) > r {
		p = p[:r]
	}

	n, err = d.Source.Read(p)
	d.n += n

	for i := 0; i < n; i++ {
		d.Key, p[i] = p[i], p[i]^d.Key
	}

	if err == io.EOF && (d.n-encSzN) != d.sz {
		err = fmt.Errorf("unexpected eof")
	}
	return n, err
}
