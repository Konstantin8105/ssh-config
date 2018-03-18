package ssh_config

// GSSAPIKeyExchange
//         Specifies whether key exchange based on GSSAPI may be used.
//         When using GSSAPI key exchange the server need not have a host
//         key.  The default is “no”.
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
	ssh_init(GSSAPIKeyExchange, funcInit)
	ssh_valid(GSSAPIKeyExchange, funcValid)
}
