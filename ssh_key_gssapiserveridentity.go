package ssh_config

// GSSAPIServerIdentity
//         If set, specifies the GSSAPI server identity that ssh should
//         expect when connecting to the server. The default is unset,
//         which means that the expected GSSAPI server identity will be
//         determined from the target hostname.
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

	sshInit(GSSAPIServerIdentity, funcInit)
	sshValid(GSSAPIServerIdentity, funcValid)
	sshParse(GSSAPIServerIdentity, funcParse)
}
