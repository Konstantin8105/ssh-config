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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(ServerAliveInterval, funcInit)
	ssh_valid(ServerAliveInterval, funcValid)
}
