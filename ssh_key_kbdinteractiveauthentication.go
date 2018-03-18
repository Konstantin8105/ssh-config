package ssh_config

// KbdInteractiveAuthentication
//         Specifies whether to use keyboard-interactive authentication.
//         The argument to this keyword must be “yes” or “no”.  The
//         default is “yes”.
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

	sshInit(KbdInteractiveAuthentication, funcInit)
	sshValid(KbdInteractiveAuthentication, funcValid)
	sshParse(KbdInteractiveAuthentication, funcParse)
}
