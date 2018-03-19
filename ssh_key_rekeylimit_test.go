package ssh_config

import "testing"

// RekeyLimit
//         Specifies the maximum amount of data that may be transmitted
//         before the session key is renegotiated, optionally followed a
//         maximum amount of time that may pass before the session key is
//         renegotiated.  The first argument is specified in bytes and may
//         have a suffix of ‘K’, ‘M’, or ‘G’ to indicate Kilobytes,
//         Megabytes, or Gigabytes, respectively.  The default is between
//         ‘1G’ and ‘4G’, depending on the cipher.  The optional second
//         value is specified in seconds and may use any of the units doc‐
//         umented in the TIME FORMATS section of sshd_config(5).  The
//         default value for RekeyLimit is “default none”, which means
//         that rekeying is performed after the cipher's default amount of
//         data has been sent or received and no time based rekeying is
//         done.
//

func TestRekeyLimitInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[RekeyLimit]
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

func TestRekeyLimitValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[RekeyLimit]
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

func TestRekeyLimitParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[RekeyLimit]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
