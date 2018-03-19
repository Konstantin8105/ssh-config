package ssh_config

import "testing"

// LocalCommand
//         Specifies a command to execute on the local machine after suc‐
//         cessfully connecting to the server.  The command string extends
//         to the end of the line, and is executed with the user's shell.
//         The following escape character substitutions will be performed:
//         ‘%d’ (local user's home directory), ‘%h’ (remote host name),
//         ‘%l’ (local host name), ‘%n’ (host name as provided on the com‐
//         mand line), ‘%p’ (remote port), ‘%r’ (remote user name) or ‘%u’
//         (local user name) or ‘%C’ by a hash of the concatenation:
//         %l%h%p%r.
//
//         The command is run synchronously and does not have access to
//         the session of the ssh(1) that spawned it.  It should not be
//         used for interactive commands.
//
//         This directive is ignored unless PermitLocalCommand has been
//         enabled.
//

func TestLocalCommandInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[LocalCommand]
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

func TestLocalCommandValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[LocalCommand]
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

func TestLocalCommandParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[LocalCommand]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
