package ssh_config

// PasswordAuthentication
//         Specifies whether to use password authentication.  The argument
//         to this keyword must be “yes” or “no”.  The default is “yes”.
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
	ssh_init(PasswordAuthentication, funcInit)
	ssh_valid(PasswordAuthentication, funcValid)
}
