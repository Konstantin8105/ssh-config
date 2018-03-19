package ssh_config

import "testing"

// MACs    Specifies the MAC (message authentication code) algorithms in
//         order of preference.  The MAC algorithm is used for data
//         integrity protection.  Multiple algorithms must be comma-sepa‐
//         rated.  If the specified value begins with a ‘+’ character,
//         then the specified algorithms will be appended to the default
//         set instead of replacing them.
//
//         The algorithms that contain “-etm” calculate the MAC after
//         encryption (encrypt-then-mac).  These are considered safer and
//         their use recommended.
//
//         The default is:
//
//               umac-64-etm@openssh.com,umac-128-etm@openssh.com,
//               hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,
//               hmac-sha1-etm@openssh.com,
//               umac-64@openssh.com,umac-128@openssh.com,
//               hmac-sha2-256,hmac-sha2-512,hmac-sha1
//
//         The list of available MAC algorithms may also be obtained using
//         the -Q option of ssh(1) with an argument of “mac”.
//

func TestMACsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[MACs]
	if !ok {
		t.Errorf("Cannot found init-function")
	}
	if f == nil {
		t.Errorf("Init-function is nil")
	}
	if f() == "" {
		// TODO
		t.Errorf("Not acceptable empty init value")
	}
}

func TestMACsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[MACs]
	if !ok {
		t.Errorf("Cannot found valid-function")
	}
	if f == nil {
		t.Errorf("Valid-function is nil")
	}
	if !f("") {
		// TODO
		t.Errorf("Not acceptable empty valid-value")
	}
}

func TestMACsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[MACs]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
