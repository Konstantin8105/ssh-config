package ssh_config

// ServerAliveCountMax
//         Sets the number of server alive messages (see below) which may
//         be sent without ssh(1) receiving any messages back from the
//         server.  If this threshold is reached while server alive mes‐
//         sages are being sent, ssh will disconnect from the server, ter‐
//         minating the session.  It is important to note that the use of
//         server alive messages is very different from TCPKeepAlive
//         (below).  The server alive messages are sent through the
//         encrypted channel and therefore will not be spoofable.  The TCP
//         keepalive option enabled by TCPKeepAlive is spoofable.  The
//         server alive mechanism is valuable when the client or server
//         depend on knowing when a connection has become inactive.
//
//         The default value is 3.  If, for example, ServerAliveInterval
//         (see below) is set to 15 and ServerAliveCountMax is left at the
//         default, if the server becomes unresponsive, ssh will discon‐
//         nect after approximately 45 seconds.
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
	ssh_init(ServerAliveCountMax, funcInit)
	ssh_valid(ServerAliveCountMax, funcValid)
}
