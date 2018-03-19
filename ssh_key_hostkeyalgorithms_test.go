package ssh_config

import "testing"

// HostKeyAlgorithms
//         Specifies the host key algorithms that the client wants to use
//         in order of preference.  Alternately if the specified value
//         begins with a ‘+’ character, then the specified key types will
//         be appended to the default set instead of replacing them.  The
//         default for this option is:
//
//            ecdsa-sha2-nistp256-cert-v01@openssh.com,
//            ecdsa-sha2-nistp384-cert-v01@openssh.com,
//            ecdsa-sha2-nistp521-cert-v01@openssh.com,
//            ssh-ed25519-cert-v01@openssh.com,
//            ssh-rsa-cert-v01@openssh.com,
//            ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,
//            ssh-ed25519,ssh-rsa
//
//         If hostkeys are known for the destination host then this
//         default is modified to prefer their algorithms.
//
//         The list of available key types may also be obtained using the
//         -Q option of ssh(1) with an argument of “key”.
//

func TestHostKeyAlgorithmsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[HostKeyAlgorithms]
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

func TestHostKeyAlgorithmsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[HostKeyAlgorithms]
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

func TestHostKeyAlgorithmsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[HostKeyAlgorithms]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
