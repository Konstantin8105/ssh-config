package ssh_config

// ControlPath
//         Specify the path to the control socket used for connection
//         sharing as described in the ControlMaster section above or the
//         string “none” to disable connection sharing.  In the path, ‘%L’
//         will be substituted by the first component of the local host
//         name, ‘%l’ will be substituted by the local host name (includ‐
//         ing any domain name), ‘%h’ will be substituted by the target
//         host name, ‘%n’ will be substituted by the original target host
//         name specified on the command line, ‘%p’ the destination port,
//         ‘%r’ by the remote login username, ‘%u’ by the username and
//         ‘%i’ by the numeric user ID (uid) of the user running ssh(1),
//         and ‘%C’ by a hash of the concatenation: %l%h%p%r.  It is rec‐
//         ommended that any ControlPath used for opportunistic connection
//         sharing include at least %h, %p, and %r (or alternatively %C)
//         and be placed in a directory that is not writable by other
//         users.  This ensures that shared connections are uniquely iden‐
//         tified.
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

	sshInit(ControlPath, funcInit)
	sshValid(ControlPath, funcValid)
	sshParse(ControlPath, funcParse)
}
