package ssh_config

// GlobalKnownHostsFile
//         Specifies one or more files to use for the global host key
//         database, separated by whitespace.  The default is
//         /etc/ssh/ssh_known_hosts, /etc/ssh/ssh_known_hosts2.
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

	sshInit(GlobalKnownHostsFile, funcInit)
	sshValid(GlobalKnownHostsFile, funcValid)
	sshParse(GlobalKnownHostsFile, funcParse)
}
