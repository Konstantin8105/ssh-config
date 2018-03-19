package ssh_config

import "testing"

// HostbasedKeyTypes
//         Specifies the key types that will be used for hostbased authen‐
//         tication as a comma-separated pattern list.  Alternately if the
//         specified value begins with a ‘+’ character, then the specified
//         key types will be appended to the default set instead of
//         replacing them.  The default for this option is:
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

func TestHostbasedKeyTypesInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[HostbasedKeyTypes]
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

func TestHostbasedKeyTypesValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[HostbasedKeyTypes]
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

func TestHostbasedKeyTypesParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[HostbasedKeyTypes]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
