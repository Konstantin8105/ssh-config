package ssh_config

// Compression
//         Specifies whether to use compression.  The argument must be
//         “yes” or “no”.  The default is “no”.
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

	sshInit(Compression, funcInit)
	sshValid(Compression, funcValid)
	sshParse(Compression, funcParse)
}
