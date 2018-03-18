package ssh_config

// GSSAPIDelegateCredentials
//         Forward (delegate) credentials to the server.  The default is
//         “no”.
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

	sshInit(GSSAPIDelegateCredentials, funcInit)
	sshValid(GSSAPIDelegateCredentials, funcValid)
	sshParse(GSSAPIDelegateCredentials, funcParse)
}
