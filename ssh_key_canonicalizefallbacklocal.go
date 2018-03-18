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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(CanonicalizeFallbackLocal, funcInit)
	ssh_valid(CanonicalizeFallbackLocal, funcValid)
}
