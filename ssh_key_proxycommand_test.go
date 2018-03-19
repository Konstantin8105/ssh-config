package ssh_config

import "testing"

// ProxyCommand
//         Specifies the command to use to connect to the server.  The
//         command string extends to the end of the line, and is executed
//         using the user's shell ‘exec’ directive to avoid a lingering
//         shell process.
//
//         In the command string, any occurrence of ‘%h’ will be substi‐
//         tuted by the host name to connect, ‘%p’ by the port, and ‘%r’
//         by the remote user name.  The command can be basically any‐
//         thing, and should read from its standard input and write to its
//         standard output.  It should eventually connect an sshd(8)
//         server running on some machine, or execute sshd -i somewhere.
//         Host key management will be done using the HostName of the host
//         being connected (defaulting to the name typed by the user).
//         Setting the command to “none” disables this option entirely.
//         Note that CheckHostIP is not available for connects with a
//         proxy command.
//
//         This directive is useful in conjunction with nc(1) and its
//         proxy support.  For example, the following directive would con‐
//         nect via an HTTP proxy at 192.0.2.0:
//
//            ProxyCommand /usr/bin/nc -X connect -x 192.0.2.0:8080 %h %p
//

func TestProxyCommandInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ProxyCommand]
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

func TestProxyCommandValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ProxyCommand]
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

func TestProxyCommandParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ProxyCommand]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
