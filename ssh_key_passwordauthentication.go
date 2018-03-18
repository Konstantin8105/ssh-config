package ssh_config

// PasswordAuthentication
//         Specifies whether to use password authentication.  The argument
//         to this keyword must be “yes” or “no”.  The default is “yes”.
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

	sshInit(PasswordAuthentication, funcInit)
	sshValid(PasswordAuthentication, funcValid)
	sshParse(PasswordAuthentication, funcParse)
}
