package ssh_config

import "testing"

// SendEnv
//         Specifies what variables from the local environ(7) should be
//         sent to the server.  The server must also support it, and the
//         server must be configured to accept these environment vari‐
//         ables.  Note that the TERM environment variable is always sent
//         whenever a pseudo-terminal is requested as it is required by
//         the protocol.  Refer to AcceptEnv in sshd_config(5) for how to
//         configure the server.  Variables are specified by name, which
//         may contain wildcard characters.  Multiple environment vari‐
//         ables may be separated by whitespace or spread across multiple
//         SendEnv directives.  The default is not to send any environment
//         variables.
//
//         See PATTERNS for more information on patterns.
//

func TestSendEnvInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[SendEnv]
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

func TestSendEnvValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[SendEnv]
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

func TestSendEnvParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[SendEnv]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
