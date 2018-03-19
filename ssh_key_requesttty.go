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
		panic("Not implemented")
		return
	}

	funcValid := func(value string) (res bool) {
		// TODO
		panic("Not implemented")
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		panic("Not implemented")
		return
	}

	sshInit(RequestTTY, funcInit)
	sshValid(RequestTTY, funcValid)
	sshParse(RequestTTY, funcParse)
}
