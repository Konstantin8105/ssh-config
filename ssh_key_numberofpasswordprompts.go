package ssh_config

// NumberOfPasswordPrompts
//         Specifies the number of password prompts before giving up.  The
//         argument to this keyword must be an integer.  The default is 3.
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

	sshInit(NumberOfPasswordPrompts, funcInit)
	sshValid(NumberOfPasswordPrompts, funcValid)
	sshParse(NumberOfPasswordPrompts, funcParse)
}
