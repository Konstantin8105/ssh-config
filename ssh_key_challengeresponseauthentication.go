package ssh_config

// ChallengeResponseAuthentication
//         Specifies whether to use challenge-response authentication.
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

	sshInit(ChallengeResponseAuthentication, funcInit)
	sshValid(ChallengeResponseAuthentication, funcValid)
	sshParse(ChallengeResponseAuthentication, funcParse)
}
