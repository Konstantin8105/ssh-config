package ssh_config

import "testing"

// CertificateFile
//         Specifies a file from which the user's certificate is read.  A
//         corresponding private key must be provided separately in order
//         to use this certificate either from an IdentityFile directive
//         or -i flag to ssh(1), via ssh-agent(1), or via a
//         PKCS11Provider.
//
//         The file name may use the tilde syntax to refer to a user's
//         home directory or one of the following escape characters: ‘%d’
//         (local user's home directory), ‘%u’ (local user name), ‘%l’
//         (local host name), ‘%h’ (remote host name) or ‘%r’ (remote user
//         name).
//
//         It is possible to have multiple certificate files specified in
//         configuration files; these certificates will be tried in
//         sequence.  Multiple CertificateFile directives will add to the
//         list of certificates used for authentication.
//

func TestCertificateFileInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[CertificateFile]
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

func TestCertificateFileValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[CertificateFile]
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

func TestCertificateFileParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[CertificateFile]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
