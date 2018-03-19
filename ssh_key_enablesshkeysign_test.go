package ssh_config

import "testing"

// EnableSSHKeysign
//         Setting this option to “yes” in the global client configuration
//         file /etc/ssh/ssh_config enables the use of the helper program
//         ssh-keysign(8) during HostbasedAuthentication.  The argument
//         must be “yes” or “no”.  The default is “no”.  This option
//         should be placed in the non-hostspecific section.  See
//         ssh-keysign(8) for more information.
//

func TestEnableSSHKeysignInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[EnableSSHKeysign]
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

func TestEnableSSHKeysignValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[EnableSSHKeysign]
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

func TestEnableSSHKeysignParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[EnableSSHKeysign]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
