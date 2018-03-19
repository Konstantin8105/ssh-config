package ssh_config

// FingerprintHash
//         Specifies the hash algorithm used when displaying key finger‐
//         prints.  Valid options are: “md5” and “sha256”.  The default is
//         “sha256”.
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

	sshInit(FingerprintHash, funcInit)
	sshValid(FingerprintHash, funcValid)
	sshParse(FingerprintHash, funcParse)
}
