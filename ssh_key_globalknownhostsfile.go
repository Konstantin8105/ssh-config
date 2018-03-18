package ssh_config

// GlobalKnownHostsFile
//         Specifies one or more files to use for the global host key
//         database, separated by whitespace.  The default is
//         /etc/ssh/ssh_known_hosts, /etc/ssh/ssh_known_hosts2.
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

	sshInit(GlobalKnownHostsFile, funcInit)
	sshValid(GlobalKnownHostsFile, funcValid)
	sshParse(GlobalKnownHostsFile, funcParse)
}
