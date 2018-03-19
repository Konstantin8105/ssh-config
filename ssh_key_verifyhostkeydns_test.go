package ssh_config

import "testing"

// VerifyHostKeyDNS
//         Specifies whether to verify the remote key using DNS and SSHFP
//         resource records.  If this option is set to “yes”, the client
//         will implicitly trust keys that match a secure fingerprint from
//         DNS.  Insecure fingerprints will be handled as if this option
//         was set to “ask”.  If this option is set to “ask”, information
//         on fingerprint match will be displayed, but the user will still
//         need to confirm new host keys according to the
//         StrictHostKeyChecking option.  The argument must be “yes”,
//         “no”, or “ask”.  The default is “no”.
//
//         See also VERIFYING HOST KEYS in ssh(1).
//

func TestVerifyHostKeyDNSInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[VerifyHostKeyDNS]
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

func TestVerifyHostKeyDNSValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[VerifyHostKeyDNS]
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

func TestVerifyHostKeyDNSParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[VerifyHostKeyDNS]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
