package ssh_config

// RemoteForward
//         Specifies that a TCP port on the remote machine be forwarded
//         over the secure channel to the specified host and port from the
//         local machine.  The first argument must be [bind_address:]port
//         and the second argument must be host:hostport.  IPv6 addresses
//         can be specified by enclosing addresses in square brackets.
//         Multiple forwardings may be specified, and additional forward‐
//         ings can be given on the command line.  Privileged ports can be
//         forwarded only when logging in as root on the remote machine.
//
//         If the port argument is ‘0’, the listen port will be dynami‐
//         cally allocated on the server and reported to the client at run
//         time.
//
//         If the bind_address is not specified, the default is to only
//         bind to loopback addresses.  If the bind_address is ‘*’ or an
//         empty string, then the forwarding is requested to listen on all
//         interfaces.  Specifying a remote bind_address will only succeed
//         if the server's GatewayPorts option is enabled (see
//         sshd_config(5)).
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

	sshInit(RemoteForward, funcInit)
	sshValid(RemoteForward, funcValid)
	sshParse(RemoteForward, funcParse)
}
