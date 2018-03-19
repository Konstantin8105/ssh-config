package ssh_config

import "testing"

// IdentitiesOnly
//         Specifies that ssh(1) should only use the authentication iden‐
//         tity and certificate files explicitly configured in the
//         ssh_config files or passed on the ssh(1) command-line, even if
//         ssh-agent(1) or a PKCS11Provider offers more identities.  The
//         argument to this keyword must be “yes” or “no”.  This option is
//         intended for situations where ssh-agent offers many different
//         identities.  The default is “no”.
//

func TestIdentitiesOnlyInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[IdentitiesOnly]
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

func TestIdentitiesOnlyValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[IdentitiesOnly]
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

func TestIdentitiesOnlyParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[IdentitiesOnly]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
