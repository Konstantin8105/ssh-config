package ssh_config

import "testing"

// HostName
//         Specifies the real host name to log into.  This can be used to
//         specify nicknames or abbreviations for hosts.  If the hostname
//         contains the character sequence ‘%h’, then this will be
//         replaced with the host name specified on the command line (this
//         is useful for manipulating unqualified names).  The character
//         sequence ‘%%’ will be replaced by a single ‘%’ character, which
//         may be used when specifying IPv6 link-local addresses.
//
//         The default is the name given on the command line.  Numeric IP
//         addresses are also permitted (both on the command line and in
//         HostName specifications).
//

func TestHostNameInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[HostName]
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

func TestHostNameValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[HostName]
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

func TestHostNameParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[HostName]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
