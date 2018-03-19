package ssh_config

import "testing"

// PubkeyAcceptedKeyTypes
//         Specifies the key types that will be used for public key
//         authentication as a comma-separated pattern list.  Alternately
//         if the specified value begins with a ‘+’ character, then the
//         key types after it will be appended to the default instead of
//         replacing it.  The default for this option is:
//
//            ecdsa-sha2-nistp256-cert-v01@openssh.com,
//            ecdsa-sha2-nistp384-cert-v01@openssh.com,
//            ecdsa-sha2-nistp521-cert-v01@openssh.com,
//            ssh-ed25519-cert-v01@openssh.com,
//            ssh-rsa-cert-v01@openssh.com,
//            ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,
//            ssh-ed25519,ssh-rsa
//
//         The -Q option of ssh(1) may be used to list supported key
//         types.
//

func TestPubkeyAcceptedKeyTypesInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[PubkeyAcceptedKeyTypes]
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

func TestPubkeyAcceptedKeyTypesValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[PubkeyAcceptedKeyTypes]
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

func TestPubkeyAcceptedKeyTypesParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[PubkeyAcceptedKeyTypes]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
