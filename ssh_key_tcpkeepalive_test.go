package ssh_config

import "testing"

// TCPKeepAlive
//         Specifies whether the system should send TCP keepalive messages
//         to the other side.  If they are sent, death of the connection
//         or crash of one of the machines will be properly noticed.  This
//         option only uses TCP keepalives (as opposed to using ssh level
//         keepalives), so takes a long time to notice when the connection
//         dies.  As such, you probably want the ServerAliveInterval
//         option as well.  However, this means that connections will die
//         if the route is down temporarily, and some people find it
//         annoying.
//
//         The default is “yes” (to send TCP keepalive messages), and the
//         client will notice if the network goes down or the remote host
//         dies.  This is important in scripts, and many users want it
//         too.
//
//         To disable TCP keepalive messages, the value should be set to
//         “no”.
//

func TestTCPKeepAliveInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[TCPKeepAlive]
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

func TestTCPKeepAliveValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[TCPKeepAlive]
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

func TestTCPKeepAliveParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[TCPKeepAlive]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
