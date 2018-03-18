package ssh_config

// CanonicalDomains
//         When CanonicalizeHostname is enabled, this option specifies the
//         list of domain suffixes in which to search for the specified
//         destination host.
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
	ssh_init(CanonicalDomains, funcInit)
	ssh_valid(CanonicalDomains, funcValid)
}
