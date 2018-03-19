package ssh_config

import "testing"

// IPQoS   Specifies the IPv4 type-of-service or DSCP class for connec‐
//         tions.  Accepted values are “af11”, “af12”, “af13”, “af21”,
//         “af22”, “af23”, “af31”, “af32”, “af33”, “af41”, “af42”, “af43”,
//         “cs0”, “cs1”, “cs2”, “cs3”, “cs4”, “cs5”, “cs6”, “cs7”, “ef”,
//         “lowdelay”, “throughput”, “reliability”, or a numeric value.
//         This option may take one or two arguments, separated by white‐
//         space.  If one argument is specified, it is used as the packet
//         class unconditionally.  If two values are specified, the first
//         is automatically selected for interactive sessions and the sec‐
//         ond for non-interactive sessions.  The default is “lowdelay”
//         for interactive sessions and “throughput” for non-interactive
//         sessions.
//

func TestIPQoSInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[IPQoS]
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

func TestIPQoSValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[IPQoS]
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

func TestIPQoSParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[IPQoS]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
