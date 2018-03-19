package ssh_config

import "testing"

// ClearAllForwardings
//         Specifies that all local, remote, and dynamic port forwardings
//         specified in the configuration files or on the command line be
//         cleared.  This option is primarily useful when used from the
//         ssh(1) command line to clear port forwardings set in configura‐
//         tion files, and is automatically set by scp(1) and sftp(1).
//         The argument must be “yes” or “no”.  The default is “no”.
//

func TestClearAllForwardingsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ClearAllForwardings]
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

func TestClearAllForwardingsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ClearAllForwardings]
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

func TestClearAllForwardingsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ClearAllForwardings]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
