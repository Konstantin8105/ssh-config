package ssh_config

// KbdInteractiveAuthentication
//         Specifies whether to use keyboard-interactive authentication.
//         The argument to this keyword must be “yes” or “no”.  The
//         default is “yes”.
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

	sshInit(KbdInteractiveAuthentication, funcInit)
	sshValid(KbdInteractiveAuthentication, funcValid)
	sshParse(KbdInteractiveAuthentication, funcParse)
}
