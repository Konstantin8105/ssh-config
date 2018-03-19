package ssh_config

import "testing"

// StreamLocalBindMask
//         Sets the octal file creation mode mask (umask) used when creat‐
//         ing a Unix-domain socket file for local or remote port forward‐
//         ing.  This option is only used for port forwarding to a Unix-
//         domain socket file.
//
//         The default value is 0177, which creates a Unix-domain socket
//         file that is readable and writable only by the owner.  Note
//         that not all operating systems honor the file mode on Unix-
//         domain socket files.
//

func TestStreamLocalBindMaskInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[StreamLocalBindMask]
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

func TestStreamLocalBindMaskValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[StreamLocalBindMask]
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

func TestStreamLocalBindMaskParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[StreamLocalBindMask]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
