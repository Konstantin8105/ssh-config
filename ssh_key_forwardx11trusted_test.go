package ssh_config

import "testing"

// ForwardX11Trusted
//         If this option is set to “yes”, remote X11 clients will have
//         full access to the original X11 display.
//
//         If this option is set to “no”, remote X11 clients will be con‐
//         sidered untrusted and prevented from stealing or tampering with
//         data belonging to trusted X11 clients.  Furthermore, the
//         xauth(1) token used for the session will be set to expire after
//         20 minutes.  Remote clients will be refused access after this
//         time.
//
//         The default is “yes” (Debian-specific).
//
//         See the X11 SECURITY extension specification for full details
//         on the restrictions imposed on untrusted clients.
//

func TestForwardX11TrustedInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ForwardX11Trusted]
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

func TestForwardX11TrustedValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ForwardX11Trusted]
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

func TestForwardX11TrustedParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ForwardX11Trusted]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
