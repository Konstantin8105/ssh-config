package ssh_config

import "testing"

// ForwardAgent
//         Specifies whether the connection to the authentication agent
//         (if any) will be forwarded to the remote machine.  The argument
//         must be “yes” or “no”.  The default is “no”.
//
//         Agent forwarding should be enabled with caution.  Users with
//         the ability to bypass file permissions on the remote host (for
//         the agent's Unix-domain socket) can access the local agent
//         through the forwarded connection.  An attacker cannot obtain
//         key material from the agent, however they can perform opera‐
//         tions on the keys that enable them to authenticate using the
//         identities loaded into the agent.
//

func TestForwardAgentInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ForwardAgent]
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

func TestForwardAgentValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ForwardAgent]
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

func TestForwardAgentParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ForwardAgent]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
