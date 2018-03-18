package ssh_config

// GSSAPIAuthentication
//         Specifies whether user authentication based on GSSAPI is
//         allowed.  The default is “no”.
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

	sshInit(GSSAPIAuthentication, funcInit)
	sshValid(GSSAPIAuthentication, funcValid)
	sshParse(GSSAPIAuthentication, funcParse)
}
