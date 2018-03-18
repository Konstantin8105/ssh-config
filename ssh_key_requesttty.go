package ssh_config

// RequestTTY
//         Specifies whether to request a pseudo-tty for the session.  The
//         argument may be one of: “no” (never request a TTY), “yes”
//         (always request a TTY when standard input is a TTY), “force”
//         (always request a TTY) or “auto” (request a TTY when opening a
//         login session).  This option mirrors the -t and -T flags for
//         ssh(1).
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
	ssh_init(RequestTTY, funcInit)
	ssh_valid(RequestTTY, funcValid)
}
