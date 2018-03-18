package ssh_config

// CanonicalizeMaxDots
//         Specifies the maximum number of dot characters in a hostname
//         before canonicalization is disabled.  The default, “1”, allows
//         a single dot (i.e. hostname.subdomain).
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
	ssh_init(CanonicalizeMaxDots, funcInit)
	ssh_valid(CanonicalizeMaxDots, funcValid)
}
