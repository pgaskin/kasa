package transport

type Requester interface {
	// Request makes a JSON request. If out is nil, optimizations may be made
	// since the response is not used. Protocol-level errors (i.e. returned in
	// the JSON) are not checked.
	Request(in, out interface{}) error
}
