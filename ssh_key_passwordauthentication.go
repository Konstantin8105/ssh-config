package ssh_config

// PasswordAuthentication
//         Specifies whether to use password authentication.  The argument
//         to this keyword must be “yes” or “no”.  The default is “yes”.
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

	sshInit(PasswordAuthentication, funcInit)
	sshValid(PasswordAuthentication, funcValid)
	sshParse(PasswordAuthentication, funcParse)
}
