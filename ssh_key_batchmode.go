package ssh_config

// BatchMode
//         If set to “yes”, passphrase/password querying will be disabled.
//         In addition, the ServerAliveInterval option will be set to 300
//         seconds by default.  This option is useful in scripts and other
//         batch jobs where no user is present to supply the password, and
//         where it is desirable to detect a broken network swiftly.  The
//         argument must be “yes” or “no”.  The default is “no”.
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

	sshInit(BatchMode, funcInit)
	sshValid(BatchMode, funcValid)
	sshParse(BatchMode, funcParse)
}
