package ssh_config

import "testing"

// RemoteForward
//         Specifies that a TCP port on the remote machine be forwarded
//         over the secure channel to the specified host and port from the
//         local machine.  The first argument must be [bind_address:]port
//         and the second argument must be host:hostport.  IPv6 addresses
//         can be specified by enclosing addresses in square brackets.
//         Multiple forwardings may be specified, and additional forward‐
//         ings can be given on the command line.  Privileged ports can be
//         forwarded only when logging in as root on the remote machine.
//
//         If the port argument is ‘0’, the listen port will be dynami‐
//         cally allocated on the server and reported to the client at run
//         time.
//
//         If the bind_address is not specified, the default is to only
//         bind to loopback addresses.  If the bind_address is ‘*’ or an
//         empty string, then the forwarding is requested to listen on all
//         interfaces.  Specifying a remote bind_address will only succeed
//         if the server's GatewayPorts option is enabled (see
//         sshd_config(5)).
//

func TestRemoteForwardInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[RemoteForward]
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

func TestRemoteForwardValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[RemoteForward]
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

func TestRemoteForwardParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[RemoteForward]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
