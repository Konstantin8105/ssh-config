package ssh_config

// Port    Specifies the port number to connect on the remote host.  The
//         default is 22.
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

	sshInit(Port, funcInit)
	sshValid(Port, funcValid)
	sshParse(Port, funcParse)
}
