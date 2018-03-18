package ssh_config

// ConnectTimeout
//         Specifies the timeout (in seconds) used when connecting to the
//         SSH server, instead of using the default system TCP timeout.
//         This value is used only when the target is down or really
//         unreachable, not when it refuses the connection.
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

	sshInit(ConnectTimeout, funcInit)
	sshValid(ConnectTimeout, funcValid)
	sshParse(ConnectTimeout, funcParse)
}
