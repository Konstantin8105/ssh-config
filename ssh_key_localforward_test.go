package ssh_config

import "testing"

// LocalForward
//         Specifies that a TCP port on the local machine be forwarded
//         over the secure channel to the specified host and port from the
//         remote machine.  The first argument must be [bind_address:]port
//         and the second argument must be host:hostport.  IPv6 addresses
//         can be specified by enclosing addresses in square brackets.
//         Multiple forwardings may be specified, and additional forward‐
//         ings can be given on the command line.  Only the superuser can
//         forward privileged ports.  By default, the local port is bound
//         in accordance with the GatewayPorts setting.  However, an
//         explicit bind_address may be used to bind the connection to a
//         specific address.  The bind_address of “localhost” indicates
//         that the listening port be bound for local use only, while an
//         empty address or ‘*’ indicates that the port should be avail‐
//         able from all interfaces.
//

func TestLocalForwardInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[LocalForward]
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

func TestLocalForwardValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[LocalForward]
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

func TestLocalForwardParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[LocalForward]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
