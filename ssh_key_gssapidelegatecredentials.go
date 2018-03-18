package ssh_config

// GSSAPIDelegateCredentials
//         Forward (delegate) credentials to the server.  The default is
//         “no”.
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
	ssh_init(GSSAPIDelegateCredentials, funcInit)
	ssh_valid(GSSAPIDelegateCredentials, funcValid)
}
