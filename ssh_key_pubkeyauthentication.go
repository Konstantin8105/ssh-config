package ssh_config

// PubkeyAuthentication
//         Specifies whether to try public key authentication.  The argu‐
//         ment to this keyword must be “yes” or “no”.  The default is
//         “yes”.
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

	sshInit(PubkeyAuthentication, funcInit)
	sshValid(PubkeyAuthentication, funcValid)
	sshParse(PubkeyAuthentication, funcParse)
}
