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
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(ChallengeResponseAuthentication, funcInit)
	ssh_valid(ChallengeResponseAuthentication, funcValid)
}
