package ssh_config

// AddressFamily
//         Specifies which address family to use when connecting.  Valid
//         arguments are “any”, “inet” (use IPv4 only), or “inet6” (use
//         IPv6 only).  The default is “any”.
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
	ssh_init(AddressFamily, funcInit)
	ssh_valid(AddressFamily, funcValid)
}
