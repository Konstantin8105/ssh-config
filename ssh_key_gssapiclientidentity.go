package ssh_config

// GSSAPIClientIdentity
//         If set, specifies the GSSAPI client identity that ssh should
//         use when connecting to the server. The default is unset, which
//         means that the default identity will be used.
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

	sshInit(GSSAPIClientIdentity, funcInit)
	sshValid(GSSAPIClientIdentity, funcValid)
	sshParse(GSSAPIClientIdentity, funcParse)
}
