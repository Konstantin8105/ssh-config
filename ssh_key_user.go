package ssh_config

// User    Specifies the user to log in as.  This can be useful when a
//         different user name is used on different machines.  This saves
//         the trouble of having to remember to give the user name on the
//         command line.
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

	sshInit(User, funcInit)
	sshValid(User, funcValid)
	sshParse(User, funcParse)
}
