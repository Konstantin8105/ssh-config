package ssh_config

import "testing"

// GatewayPorts
//         Specifies whether remote hosts are allowed to connect to local
//         forwarded ports.  By default, ssh(1) binds local port forward‐
//         ings to the loopback address.  This prevents other remote hosts
//         from connecting to forwarded ports.  GatewayPorts can be used
//         to specify that ssh should bind local port forwardings to the
//         wildcard address, thus allowing remote hosts to connect to for‐
//         warded ports.  The argument must be “yes” or “no”.  The default
//         is “no”.
//

func TestGatewayPortsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[GatewayPorts]
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

func TestGatewayPortsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[GatewayPorts]
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

func TestGatewayPortsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[GatewayPorts]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
