package ssh_config

import "testing"

// AddKeysToAgent
//         Specifies whether keys should be automatically added to a run‐
//         ning ssh-agent(1).  If this option is set to “yes” and a key is
//         loaded from a file, the key and its passphrase are added to the
//         agent with the default lifetime, as if by ssh-add(1).  If this
//         option is set to “ask”, ssh will require confirmation using the
//         SSH_ASKPASS program before adding a key (see ssh-add(1) for
//         details).  If this option is set to “confirm”, each use of the
//         key must be confirmed, as if the -c option was specified to
//         ssh-add(1).  If this option is set to “no”, no keys are added
//         to the agent.  The argument must be “yes”, “confirm”, “ask”, or
//         “no”.  The default is “no”.
//

func TestAddKeysToAgentInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[AddKeysToAgent]
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

func TestAddKeysToAgentValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[AddKeysToAgent]
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

func TestAddKeysToAgentParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[AddKeysToAgent]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
