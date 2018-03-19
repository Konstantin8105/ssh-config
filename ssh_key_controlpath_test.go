package ssh_config

import "testing"

// ControlPath
//         Specify the path to the control socket used for connection
//         sharing as described in the ControlMaster section above or the
//         string “none” to disable connection sharing.  In the path, ‘%L’
//         will be substituted by the first component of the local host
//         name, ‘%l’ will be substituted by the local host name (includ‐
//         ing any domain name), ‘%h’ will be substituted by the target
//         host name, ‘%n’ will be substituted by the original target host
//         name specified on the command line, ‘%p’ the destination port,
//         ‘%r’ by the remote login username, ‘%u’ by the username and
//         ‘%i’ by the numeric user ID (uid) of the user running ssh(1),
//         and ‘%C’ by a hash of the concatenation: %l%h%p%r.  It is rec‐
//         ommended that any ControlPath used for opportunistic connection
//         sharing include at least %h, %p, and %r (or alternatively %C)
//         and be placed in a directory that is not writable by other
//         users.  This ensures that shared connections are uniquely iden‐
//         tified.
//

func TestControlPathInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ControlPath]
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

func TestControlPathValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ControlPath]
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

func TestControlPathParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ControlPath]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
