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

	funcValid := func(value string) (res bool) {
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		return
	}

	sshInit(ConnectionAttempts, funcInit)
	sshValid(ConnectionAttempts, funcValid)
	sshParse(ConnectionAttempts, funcParse)
}
