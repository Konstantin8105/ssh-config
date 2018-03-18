package ssh_config

// VisualHostKey
//         If this flag is set to “yes”, an ASCII art representation of
//         the remote host key fingerprint is printed in addition to the
//         fingerprint string at login and for unknown host keys.  If this
//         flag is set to “no”, no fingerprint strings are printed at
//         login and only the fingerprint string will be printed for
//         unknown host keys.  The default is “no”.
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
	ssh_init(VisualHostKey, funcInit)
	ssh_valid(VisualHostKey, funcValid)
}
