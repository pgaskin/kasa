package tpcommand

import (
	"testing"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		Error       Error
		CodeString  string
		ErrorString string
	}{
		{Error{0, ""},
			"no error",
			"\x00"},
		{Error{0, "test"},
			"no error",
			"tp-link api error: test"},
		{Error{1, ""},
			"no error",
			"\x00"},
		{Error{1, "test"},
			"no error",
			"\x00"},
		{Error{-20511, ""},
			"-20511 (APP_INVALID_DEVICE)",
			"tp-link api error -20511 (APP_INVALID_DEVICE)"},
		{Error{-20511, "test"},
			"-20511 (APP_INVALID_DEVICE)",
			"tp-link api error -20511 (APP_INVALID_DEVICE): test"},
		{Error{-99102, ""},
			"-99102 (AUTH_CLOUD_SERVICES_NOT_INITIALIZED || DEVICE_CACHE_FAILED_TO_RETRIEVE)",
			"tp-link api error -99102 (AUTH_CLOUD_SERVICES_NOT_INITIALIZED || DEVICE_CACHE_FAILED_TO_RETRIEVE)"},
		{Error{-99102, "test"},
			"-99102 (AUTH_CLOUD_SERVICES_NOT_INITIALIZED || DEVICE_CACHE_FAILED_TO_RETRIEVE)",
			"tp-link api error -99102 (AUTH_CLOUD_SERVICES_NOT_INITIALIZED || DEVICE_CACHE_FAILED_TO_RETRIEVE): test"},
	} {
		if v := tc.Error.Code.String(); v != tc.CodeString {
			t.Errorf("expected %#v code string to be %q, got %q", tc.Error, tc.CodeString, v)
		}
		if tc.ErrorString == "\x00" {
			var panicked bool
			func() {
				defer func() {
					if recover() != nil {
						panicked = true
					}
				}()
				_ = tc.Error.Error()
			}()
			if !panicked {
				t.Errorf("expected %#v error to panic since it isn't actually an error", tc.Error)
			}
		} else if v := tc.Error.Error(); v != tc.ErrorString {
			t.Errorf("expected %#v error string to be %q, got %q", tc.Error, tc.ErrorString, v)
		}
	}
}
