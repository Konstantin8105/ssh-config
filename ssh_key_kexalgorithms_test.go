package ssh_config

import "testing"

// KexAlgorithms
//         Specifies the available KEX (Key Exchange) algorithms.  Multi‐
//         ple algorithms must be comma-separated.  Alternately if the
//         specified value begins with a ‘+’ character, then the specified
//         methods will be appended to the default set instead of replac‐
//         ing them.  The default is:
//
//               curve25519-sha256@libssh.org,
//               ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,
//               diffie-hellman-group-exchange-sha256,
//               diffie-hellman-group-exchange-sha1,
//               diffie-hellman-group14-sha1
//
//         The list of available key exchange algorithms may also be
//         obtained using the -Q option of ssh(1) with an argument of
//         “kex”.
//

func TestKexAlgorithmsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[KexAlgorithms]
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

func TestKexAlgorithmsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[KexAlgorithms]
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

func TestKexAlgorithmsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[KexAlgorithms]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
