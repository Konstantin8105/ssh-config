package ssh_config

import "testing"

// ControlMaster
//         Enables the sharing of multiple sessions over a single network
//         connection.  When set to “yes”, ssh(1) will listen for connec‐
//         tions on a control socket specified using the ControlPath argu‐
//         ment.  Additional sessions can connect to this socket using the
//         same ControlPath with ControlMaster set to “no” (the default).
//         These sessions will try to reuse the master instance's network
//         connection rather than initiating new ones, but will fall back
//         to connecting normally if the control socket does not exist, or
//         is not listening.
//
//         Setting this to “ask” will cause ssh to listen for control con‐
//         nections, but require confirmation using ssh-askpass(1).  If
//         the ControlPath cannot be opened, ssh will continue without
//         connecting to a master instance.
//
//         X11 and ssh-agent(1) forwarding is supported over these multi‐
//         plexed connections, however the display and agent forwarded
//         will be the one belonging to the master connection i.e. it is
//         not possible to forward multiple displays or agents.
//
//         Two additional options allow for opportunistic multiplexing:
//         try to use a master connection but fall back to creating a new
//         one if one does not already exist.  These options are: “auto”
//         and “autoask”.  The latter requires confirmation like the “ask”
//         option.
//

func TestControlMasterInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ControlMaster]
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

func TestControlMasterValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ControlMaster]
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

func TestControlMasterParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ControlMaster]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
