package ssh_config

import "testing"

// NoHostAuthenticationForLocalhost
//         This option can be used if the home directory is shared across
//         machines.  In this case localhost will refer to a different
//         machine on each of the machines and the user will get many
//         warnings about changed host keys.  However, this option dis‐
//         ables host authentication for localhost.  The argument to this
//         keyword must be “yes” or “no”.  The default is to check the
//         host key for localhost.
//

func TestNoHostAuthenticationForLocalhostInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[NoHostAuthenticationForLocalhost]
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

func TestNoHostAuthenticationForLocalhostValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[NoHostAuthenticationForLocalhost]
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

func TestNoHostAuthenticationForLocalhostParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[NoHostAuthenticationForLocalhost]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
