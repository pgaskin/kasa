package tpcommand

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestChecked(t *testing.T) {
	tcs := []struct {
		JSON  string
		Error error
	}{
		{``, nil},
		{`"err_code": 0, "err_msg": "", `, nil},
		{`"err_code": 0, "err_msg": "test", `, Error{Code: 0, Message: "test"}},
		{`"err_code": -1, "err_msg": "test", `, Error{Code: -1, Message: "test"}},
		{`"err_code": -1, "err_msg": "", `, Error{Code: -1, Message: ""}},
	}
	t.Run("Direct", func(t *testing.T) {
		for _, tc := range tcs {
			var c Checked
			str := `{` + strings.TrimSuffix(tc.JSON, ", ") + `}`
			if err := json.Unmarshal([]byte(str), &c); err != nil {
				t.Errorf("unmarshal `%s` to %T: %v", str, c, err)
			} else if err := c.CheckError(); err != tc.Error {
				t.Errorf("unmarshal `%s` to %T: expected %#v, got %#v", str, c, tc.Error, err)
			}
		}
	})
	t.Run("Embedded", func(t *testing.T) {
		for _, tc := range tcs {
			var c struct {
				Other struct {
					String string
				}
				Checked
			}
			str := `{` + tc.JSON + `"other": {"string": "test"}}`
			if err := json.Unmarshal([]byte(str), &c); err != nil {
				t.Errorf("unmarshal `%s` to %T: %v", str, c, err)
			} else if err := c.CheckError(); err != tc.Error {
				t.Errorf("unmarshal `%s` to %T: expected %#v, got %#v", str, c, tc.Error, err)
			} else if c.Other.String != "test" {
				t.Errorf("unmarshal `%s` to %T: expected other fields to be correct", str, c)
			}
		}
	})
	t.Run("EmbeddedNested", func(t *testing.T) {
		for _, tc := range tcs {
			type n struct {
				Other1 struct {
					String string
				}
				Checked
			}
			var c struct {
				Other struct {
					String string
				}
				n
			}
			str := `{` + tc.JSON + `"other": {"string": "test"}, "other1": {"string": "test1"}}`
			if err := json.Unmarshal([]byte(str), &c); err != nil {
				t.Errorf("unmarshal `%s` to %T: %v", str, c, err)
			} else if err := c.CheckError(); err != tc.Error {
				t.Errorf("unmarshal `%s` to %T: expected %#v, got %#v", str, c, tc.Error, err)
			} else if c.Other.String != "test" || c.Other1.String != "test1" {
				t.Errorf("unmarshal `%s` to %T: expected other fields to be correct", str, c)
			}
		}
	})
	t.Run("Complex", func(t *testing.T) {
		for _, tc := range tcs {
			type n struct {
				Other1 struct {
					String string
					Checked
					Test int
				}
				Checked
			}
			var c struct {
				Other struct {
					String string
				}
				n
			}
			str := `{` + tc.JSON + `"other": {"string": "test"}, "other1": {` + tc.JSON + `"string": "test1", "test": 1}}`
			if err := json.Unmarshal([]byte(str), &c); err != nil {
				t.Errorf("unmarshal `%s` to %T: %v", str, c, err)
			} else if err := c.CheckError(); err != tc.Error {
				t.Errorf("unmarshal `%s` to %T: expected %#v, got %#v", str, c, tc.Error, err)
			} else if err := c.Other1.CheckError(); err != tc.Error {
				t.Errorf("unmarshal `%s` to %T: expected %#v, got %#v", str, c, tc.Error, err)
			} else if c.Other.String != "test" || c.Other1.String != "test1" || c.Other1.Test != 1 {
				t.Errorf("unmarshal `%s` to %T: expected other fields to be correct", str, c)
			}
		}
	})
}
