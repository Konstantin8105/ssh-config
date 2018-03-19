package ssh_config

import "testing"

// UpdateHostKeys
//         Specifies whether ssh(1) should accept notifications of addi‐
//         tional hostkeys from the server sent after authentication has
//         completed and add them to UserKnownHostsFile.  The argument
//         must be “yes”, “no” (the default) or “ask”.  Enabling this
//         option allows learning alternate hostkeys for a server and sup‐
//         ports graceful key rotation by allowing a server to send
//         replacement public keys before old ones are removed.  Addi‐
//         tional hostkeys are only accepted if the key used to authenti‐
//         cate the host was already trusted or explicitly accepted by the
//         user.  If UpdateHostKeys is set to “ask”, then the user is
//         asked to confirm the modifications to the known_hosts file.
//         Confirmation is currently incompatible with ControlPersist, and
//         will be disabled if it is enabled.
//
//         Presently, only sshd(8) from OpenSSH 6.8 and greater support
//         the “hostkeys@openssh.com” protocol extension used to inform
//         the client of all the server's hostkeys.
//

func TestUpdateHostKeysInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[UpdateHostKeys]
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

func TestUpdateHostKeysValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[UpdateHostKeys]
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

func TestUpdateHostKeysParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[UpdateHostKeys]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
