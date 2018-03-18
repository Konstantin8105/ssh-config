package ssh_config

// GSSAPIClientIdentity
//         If set, specifies the GSSAPI client identity that ssh should
//         use when connecting to the server. The default is unset, which
//         means that the default identity will be used.
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
	ssh_init(GSSAPIClientIdentity, funcInit)
	ssh_valid(GSSAPIClientIdentity, funcValid)
}
