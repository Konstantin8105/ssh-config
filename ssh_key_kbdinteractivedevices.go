package ssh_config

// KbdInteractiveDevices
//         Specifies the list of methods to use in keyboard-interactive
//         authentication.  Multiple method names must be comma-separated.
//         The default is to use the server specified list.  The methods
//         available vary depending on what the server supports.  For an
//         OpenSSH server, it may be zero or more of: “bsdauth”, “pam”,
//         and “skey”.
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
	ssh_init(KbdInteractiveDevices, funcInit)
	ssh_valid(KbdInteractiveDevices, funcValid)
}
