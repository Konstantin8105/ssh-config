package ssh_config

// KbdInteractiveAuthentication
//         Specifies whether to use keyboard-interactive authentication.
//         The argument to this keyword must be “yes” or “no”.  The
//         default is “yes”.
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
	ssh_init(KbdInteractiveAuthentication, funcInit)
	ssh_valid(KbdInteractiveAuthentication, funcValid)
}
