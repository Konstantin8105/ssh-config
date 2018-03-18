package ssh_config

// Compression
//         Specifies whether to use compression.  The argument must be
//         “yes” or “no”.  The default is “no”.
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

	sshInit(Compression, funcInit)
	sshValid(Compression, funcValid)
	sshParse(Compression, funcParse)
}
