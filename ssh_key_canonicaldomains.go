package ssh_config

// CanonicalDomains
//         When CanonicalizeHostname is enabled, this option specifies the
//         list of domain suffixes in which to search for the specified
//         destination host.
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

	sshInit(CanonicalDomains, funcInit)
	sshValid(CanonicalDomains, funcValid)
	sshParse(CanonicalDomains, funcParse)
}
