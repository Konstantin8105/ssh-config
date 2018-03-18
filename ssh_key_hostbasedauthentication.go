package ssh_config

// HostbasedAuthentication
//         Specifies whether to try rhosts based authentication with pub‐
//         lic key authentication.  The argument must be “yes” or “no”.
//         The default is “no”.
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

	sshInit(HostbasedAuthentication, funcInit)
	sshValid(HostbasedAuthentication, funcValid)
	sshParse(HostbasedAuthentication, funcParse)
}
