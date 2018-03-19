package ssh_config

// CanonicalizeFallbackLocal
//         Specifies whether to fail with an error when hostname canoni‐
//         calization fails.  The default, “yes”, will attempt to look up
//         the unqualified hostname using the system resolver's search
//         rules.  A value of “no” will cause ssh(1) to fail instantly if
//         CanonicalizeHostname is enabled and the target hostname cannot
//         be found in any of the domains specified by CanonicalDomains.
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

	sshInit(CanonicalizeFallbackLocal, funcInit)
	sshValid(CanonicalizeFallbackLocal, funcValid)
	sshParse(CanonicalizeFallbackLocal, funcParse)
}
