package ssh_config

// RhostsRSAAuthentication
//         Specifies whether to try rhosts based authentication with RSA
//         host authentication.  The argument must be “yes” or “no”.  The
//         default is “no”.  This option applies to protocol version 1
//         only and requires ssh(1) to be setuid root.
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
	ssh_init(RhostsRSAAuthentication, funcInit)
	ssh_valid(RhostsRSAAuthentication, funcValid)
}
