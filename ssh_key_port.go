package ssh_config

// Port    Specifies the port number to connect on the remote host.  The
//         default is 22.
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

	sshInit(Port, funcInit)
	sshValid(Port, funcValid)
	sshParse(Port, funcParse)
}
