package ssh_config

// GSSAPIAuthentication
//         Specifies whether user authentication based on GSSAPI is
//         allowed.  The default is “no”.
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
	ssh_init(GSSAPIAuthentication, funcInit)
	ssh_valid(GSSAPIAuthentication, funcValid)
}
