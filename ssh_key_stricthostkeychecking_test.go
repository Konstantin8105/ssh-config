package ssh_config

import "testing"

// StrictHostKeyChecking
//         If this flag is set to “yes”, ssh(1) will never automatically
//         add host keys to the ~/.ssh/known_hosts file, and refuses to
//         connect to hosts whose host key has changed.  This provides
//         maximum protection against trojan horse attacks, though it can
//         be annoying when the /etc/ssh/ssh_known_hosts file is poorly
//         maintained or when connections to new hosts are frequently
//         made.  This option forces the user to manually add all new
//         hosts.  If this flag is set to “no”, ssh will automatically add
//         new host keys to the user known hosts files.  If this flag is
//         set to “ask”, new host keys will be added to the user known
//         host files only after the user has confirmed that is what they
//         really want to do, and ssh will refuse to connect to hosts
//         whose host key has changed.  The host keys of known hosts will
//         be verified automatically in all cases.  The argument must be
//         “yes”, “no”, or “ask”.  The default is “ask”.
//

func TestStrictHostKeyCheckingInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[StrictHostKeyChecking]
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

func TestStrictHostKeyCheckingValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[StrictHostKeyChecking]
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

func TestStrictHostKeyCheckingParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[StrictHostKeyChecking]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
