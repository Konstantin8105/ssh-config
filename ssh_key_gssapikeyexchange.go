package ssh_config

// GSSAPIKeyExchange
//         Specifies whether key exchange based on GSSAPI may be used.
//         When using GSSAPI key exchange the server need not have a host
//         key.  The default is “no”.
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

	sshInit(GSSAPIKeyExchange, funcInit)
	sshValid(GSSAPIKeyExchange, funcValid)
	sshParse(GSSAPIKeyExchange, funcParse)
}
