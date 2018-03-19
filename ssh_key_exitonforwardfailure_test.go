package ssh_config

import "testing"

// ExitOnForwardFailure
//         Specifies whether ssh(1) should terminate the connection if it
//         cannot set up all requested dynamic, tunnel, local, and remote
//         port forwardings, (e.g. if either end is unable to bind and
//         listen on a specified port).  Note that ExitOnForwardFailure
//         does not apply to connections made over port forwardings and
//         will not, for example, cause ssh(1) to exit if TCP connections
//         to the ultimate forwarding destination fail.  The argument must
//         be “yes” or “no”.  The default is “no”.
//

func TestExitOnForwardFailureInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ExitOnForwardFailure]
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

func TestExitOnForwardFailureValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ExitOnForwardFailure]
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

func TestExitOnForwardFailureParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ExitOnForwardFailure]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
