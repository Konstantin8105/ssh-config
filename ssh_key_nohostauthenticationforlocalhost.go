package ssh_config

// NoHostAuthenticationForLocalhost
//         This option can be used if the home directory is shared across
//         machines.  In this case localhost will refer to a different
//         machine on each of the machines and the user will get many
//         warnings about changed host keys.  However, this option dis‐
//         ables host authentication for localhost.  The argument to this
//         keyword must be “yes” or “no”.  The default is to check the
//         host key for localhost.
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

	sshInit(NoHostAuthenticationForLocalhost, funcInit)
	sshValid(NoHostAuthenticationForLocalhost, funcValid)
	sshParse(NoHostAuthenticationForLocalhost, funcParse)
}
