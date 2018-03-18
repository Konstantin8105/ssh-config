package ssh_config

// ConnectionAttempts
//         Specifies the number of tries (one per second) to make before
//         exiting.  The argument must be an integer.  This may be useful
//         in scripts if the connection sometimes fails.  The default is
//         1.
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
	ssh_init(ConnectionAttempts, funcInit)
	ssh_valid(ConnectionAttempts, funcValid)
}
