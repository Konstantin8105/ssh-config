package ssh_config

// EscapeChar
//         Sets the escape character (default: ‘~’).  The escape character
//         can also be set on the command line.  The argument should be a
//         single character, ‘^’ followed by a letter, or “none” to dis‐
//         able the escape character entirely (making the connection
//         transparent for binary data).
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

	sshInit(EscapeChar, funcInit)
	sshValid(EscapeChar, funcValid)
	sshParse(EscapeChar, funcParse)
}
