package ssh_config

import "testing"

// DynamicForward
//         Specifies that a TCP port on the local machine be forwarded
//         over the secure channel, and the application protocol is then
//         used to determine where to connect to from the remote machine.
//
//         The argument must be [bind_address:]port.  IPv6 addresses can
//         be specified by enclosing addresses in square brackets.  By
//         default, the local port is bound in accordance with the
//         GatewayPorts setting.  However, an explicit bind_address may be
//         used to bind the connection to a specific address.  The
//         bind_address of “localhost” indicates that the listening port
//         be bound for local use only, while an empty address or ‘*’
//         indicates that the port should be available from all inter‐
//         faces.
//
//         Currently the SOCKS4 and SOCKS5 protocols are supported, and
//         ssh(1) will act as a SOCKS server.  Multiple forwardings may be
//         specified, and additional forwardings can be given on the com‐
//         mand line.  Only the superuser can forward privileged ports.
//

func TestDynamicForwardInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[DynamicForward]
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

func TestDynamicForwardValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[DynamicForward]
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

func TestDynamicForwardParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[DynamicForward]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
