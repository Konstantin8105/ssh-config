package ssh_config

// PubkeyAuthentication
//         Specifies whether to try public key authentication.  The argu‐
//         ment to this keyword must be “yes” or “no”.  The default is
//         “yes”.
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

	sshInit(PubkeyAuthentication, funcInit)
	sshValid(PubkeyAuthentication, funcValid)
	sshParse(PubkeyAuthentication, funcParse)
}
