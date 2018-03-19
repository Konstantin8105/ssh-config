package ssh_config

// CanonicalizeMaxDots
//         Specifies the maximum number of dot characters in a hostname
//         before canonicalization is disabled.  The default, “1”, allows
//         a single dot (i.e. hostname.subdomain).
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

	sshInit(CanonicalizeMaxDots, funcInit)
	sshValid(CanonicalizeMaxDots, funcValid)
	sshParse(CanonicalizeMaxDots, funcParse)
}
