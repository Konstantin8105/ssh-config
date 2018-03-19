package ssh_config

// ServerAliveInterval
//         Sets a timeout interval in seconds after which if no data has
//         been received from the server, ssh(1) will send a message
//         through the encrypted channel to request a response from the
//         server.  The default is 0, indicating that these messages will
//         not be sent to the server, or 300 if the BatchMode option is
//         set.  ProtocolKeepAlives and SetupTimeOut are Debian-specific
//         compatibility aliases for this option.
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

	sshInit(ServerAliveInterval, funcInit)
	sshValid(ServerAliveInterval, funcValid)
	sshParse(ServerAliveInterval, funcParse)
}
