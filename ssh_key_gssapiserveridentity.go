package ssh_config

// GSSAPIServerIdentity
//         If set, specifies the GSSAPI server identity that ssh should
//         expect when connecting to the server. The default is unset,
//         which means that the expected GSSAPI server identity will be
//         determined from the target hostname.
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
	ssh_init(GSSAPIServerIdentity, funcInit)
	ssh_valid(GSSAPIServerIdentity, funcValid)
}
