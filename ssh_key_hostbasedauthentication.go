package ssh_config

// HostbasedAuthentication
//         Specifies whether to try rhosts based authentication with pub‐
//         lic key authentication.  The argument must be “yes” or “no”.
//         The default is “no”.
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
	ssh_init(HostbasedAuthentication, funcInit)
	ssh_valid(HostbasedAuthentication, funcValid)
}
