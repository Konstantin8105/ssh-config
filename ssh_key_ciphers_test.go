package ssh_config

import "testing"

// Ciphers
//         Specifies the ciphers allowed for protocol version 2 in order
//         of preference.  Multiple ciphers must be comma-separated.  If
//         the specified value begins with a ‘+’ character, then the spec‐
//         ified ciphers will be appended to the default set instead of
//         replacing them.
//
//         The supported ciphers are:
//
//               3des-cbc
//               aes128-cbc
//               aes192-cbc
//               aes256-cbc
//               aes128-ctr
//               aes192-ctr
//               aes256-ctr
//               aes128-gcm@openssh.com
//               aes256-gcm@openssh.com
//               arcfour
//               arcfour128
//               arcfour256
//               blowfish-cbc
//               cast128-cbc
//               chacha20-poly1305@openssh.com
//
//         The default is:
//
//               chacha20-poly1305@openssh.com,
//               aes128-ctr,aes192-ctr,aes256-ctr,
//               aes128-gcm@openssh.com,aes256-gcm@openssh.com,
//               aes128-cbc,aes192-cbc,aes256-cbc,3des-cbc
//
//         The list of available ciphers may also be obtained using the -Q
//         option of ssh(1) with an argument of “cipher”.
//

func TestCiphersInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[Ciphers]
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

func TestCiphersValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[Ciphers]
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

func TestCiphersParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[Ciphers]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
