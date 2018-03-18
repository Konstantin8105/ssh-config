package ssh_config

// TCPKeepAlive
//         Specifies whether the system should send TCP keepalive messages
//         to the other side.  If they are sent, death of the connection
//         or crash of one of the machines will be properly noticed.  This
//         option only uses TCP keepalives (as opposed to using ssh level
//         keepalives), so takes a long time to notice when the connection
//         dies.  As such, you probably want the ServerAliveInterval
//         option as well.  However, this means that connections will die
//         if the route is down temporarily, and some people find it
//         annoying.
//
//         The default is “yes” (to send TCP keepalive messages), and the
//         client will notice if the network goes down or the remote host
//         dies.  This is important in scripts, and many users want it
//         too.
//
//         To disable TCP keepalive messages, the value should be set to
//         “no”.
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

	sshInit(TCPKeepAlive, funcInit)
	sshValid(TCPKeepAlive, funcValid)
	sshParse(TCPKeepAlive, funcParse)
}
