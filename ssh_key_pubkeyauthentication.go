package ssh_config

// PubkeyAuthentication
//         Specifies whether to try public key authentication.  The argu‐
//         ment to this keyword must be “yes” or “no”.  The default is
//         “yes”.
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
	ssh_init(PubkeyAuthentication, funcInit)
	ssh_valid(PubkeyAuthentication, funcValid)
}
