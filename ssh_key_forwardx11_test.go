package ssh_config

import "testing"

// ForwardX11
//         Specifies whether X11 connections will be automatically redi‐
//         rected over the secure channel and DISPLAY set.  The argument
//         must be “yes” or “no”.  The default is “no”.
//
//         X11 forwarding should be enabled with caution.  Users with the
//         ability to bypass file permissions on the remote host (for the
//         user's X11 authorization database) can access the local X11
//         display through the forwarded connection.  An attacker may then
//         be able to perform activities such as keystroke monitoring if
//         the ForwardX11Trusted option is also enabled.
//

func TestForwardX11InitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ForwardX11]
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

func TestForwardX11ValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ForwardX11]
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

func TestForwardX11ParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ForwardX11]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
