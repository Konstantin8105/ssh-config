package ssh_config

// GSSAPITrustDns
//         Set to “yes” to indicate that the DNS is trusted to securely
//         canonicalize the name of the host being connected to. If “no”,
//         the hostname entered on the command line will be passed
//         untouched to the GSSAPI library.  The default is “no”.
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
	ssh_init(GSSAPITrustDns, funcInit)
	ssh_valid(GSSAPITrustDns, funcValid)
}
