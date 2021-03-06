package ssh_config

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

func init() {
	funcInit := func() (res string) {
		// TODO
		panic("Not implemented")
		return
	}

	funcValid := func(value string) (res bool) {
		// TODO
		panic("Not implemented")
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		panic("Not implemented")
		return
	}

	sshInit(ProxyCommand, funcInit)
	sshValid(ProxyCommand, funcValid)
	sshParse(ProxyCommand, funcParse)
}
