package ssh_config

import "testing"

// StreamLocalBindUnlink
//         Specifies whether to remove an existing Unix-domain socket file
//         for local or remote port forwarding before creating a new one.
//         If the socket file already exists and StreamLocalBindUnlink is
//         not enabled, ssh will be unable to forward the port to the
//         Unix-domain socket file.  This option is only used for port
//         forwarding to a Unix-domain socket file.
//
//         The argument must be “yes” or “no”.  The default is “no”.
//

func TestStreamLocalBindUnlinkInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[StreamLocalBindUnlink]
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

func TestStreamLocalBindUnlinkValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[StreamLocalBindUnlink]
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

func TestStreamLocalBindUnlinkParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[StreamLocalBindUnlink]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
