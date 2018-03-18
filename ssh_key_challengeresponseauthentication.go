package ssh_config

// ChallengeResponseAuthentication
//         Specifies whether to use challenge-response authentication.
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

	sshInit(ChallengeResponseAuthentication, funcInit)
	sshValid(ChallengeResponseAuthentication, funcValid)
	sshParse(ChallengeResponseAuthentication, funcParse)
}
