package ssh_config

// ClearAllForwardings
//         Specifies that all local, remote, and dynamic port forwardings
//         specified in the configuration files or on the command line be
//         cleared.  This option is primarily useful when used from the
//         ssh(1) command line to clear port forwardings set in configura‐
//         tion files, and is automatically set by scp(1) and sftp(1).
//         The argument must be “yes” or “no”.  The default is “no”.
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

	sshInit(ClearAllForwardings, funcInit)
	sshValid(ClearAllForwardings, funcValid)
	sshParse(ClearAllForwardings, funcParse)
}
