package ssh_config

// GSSAPIDelegateCredentials
//         Forward (delegate) credentials to the server.  The default is
//         “no”.
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

	sshInit(GSSAPIDelegateCredentials, funcInit)
	sshValid(GSSAPIDelegateCredentials, funcValid)
	sshParse(GSSAPIDelegateCredentials, funcParse)
}
