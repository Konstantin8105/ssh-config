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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(EscapeChar, funcInit)
	ssh_valid(EscapeChar, funcValid)
}
