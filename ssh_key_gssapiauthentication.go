package ssh_config

// GSSAPIAuthentication
//         Specifies whether user authentication based on GSSAPI is
//         allowed.  The default is “no”.
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

	sshInit(GSSAPIAuthentication, funcInit)
	sshValid(GSSAPIAuthentication, funcValid)
	sshParse(GSSAPIAuthentication, funcParse)
}
