package ssh_config

// EnableSSHKeysign
//         Setting this option to “yes” in the global client configuration
//         file /etc/ssh/ssh_config enables the use of the helper program
//         ssh-keysign(8) during HostbasedAuthentication.  The argument
//         must be “yes” or “no”.  The default is “no”.  This option
//         should be placed in the non-hostspecific section.  See
//         ssh-keysign(8) for more information.
//

func init() {
	funcInit := func() (res string) {
		// TODO
		return
	}

	funcValid := func(value string) (res bool) {
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		return
	}

	sshInit(EnableSSHKeysign, funcInit)
	sshValid(EnableSSHKeysign, funcValid)
	sshParse(EnableSSHKeysign, funcParse)
}
