package ssh_config

// RhostsRSAAuthentication
//         Specifies whether to try rhosts based authentication with RSA
//         host authentication.  The argument must be “yes” or “no”.  The
//         default is “no”.  This option applies to protocol version 1
//         only and requires ssh(1) to be setuid root.
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

	sshInit(RhostsRSAAuthentication, funcInit)
	sshValid(RhostsRSAAuthentication, funcValid)
	sshParse(RhostsRSAAuthentication, funcParse)
}
