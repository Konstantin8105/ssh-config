package ssh_config

// LocalForward
//         Specifies that a TCP port on the local machine be forwarded
//         over the secure channel to the specified host and port from the
//         remote machine.  The first argument must be [bind_address:]port
//         and the second argument must be host:hostport.  IPv6 addresses
//         can be specified by enclosing addresses in square brackets.
//         Multiple forwardings may be specified, and additional forward‐
//         ings can be given on the command line.  Only the superuser can
//         forward privileged ports.  By default, the local port is bound
//         in accordance with the GatewayPorts setting.  However, an
//         explicit bind_address may be used to bind the connection to a
//         specific address.  The bind_address of “localhost” indicates
//         that the listening port be bound for local use only, while an
//         empty address or ‘*’ indicates that the port should be avail‐
//         able from all interfaces.
//

func init() {
	funcInit := func() (res string) {
		// TODO
		return
	}

	funcValid := func(value string) (res bool) {
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		return
	}

	sshInit(LocalForward, funcInit)
	sshValid(LocalForward, funcValid)
	sshParse(LocalForward, funcParse)
}
