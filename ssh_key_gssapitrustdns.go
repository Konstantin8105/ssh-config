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

	funcValid := func(value string) (res bool) {
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		return
	}

	sshInit(GSSAPITrustDns, funcInit)
	sshValid(GSSAPITrustDns, funcValid)
	sshParse(GSSAPITrustDns, funcParse)
}
