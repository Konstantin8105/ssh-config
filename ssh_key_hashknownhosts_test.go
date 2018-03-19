package ssh_config

import "testing"

// HashKnownHosts
//         Indicates that ssh(1) should hash host names and addresses when
//         they are added to ~/.ssh/known_hosts.  These hashed names may
//         be used normally by ssh(1) and sshd(8), but they do not reveal
//         identifying information should the file's contents be dis‐
//         closed.  The default is “no”.  Note that existing names and
//         addresses in known hosts files will not be converted automati‐
//         cally, but may be manually hashed using ssh-keygen(1).  Use of
//         this option may break facilities such as tab-completion that
//         rely on being able to read unhashed host names from
//         ~/.ssh/known_hosts.
//

func TestHashKnownHostsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[HashKnownHosts]
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

func TestHashKnownHostsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[HashKnownHosts]
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

func TestHashKnownHostsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[HashKnownHosts]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
