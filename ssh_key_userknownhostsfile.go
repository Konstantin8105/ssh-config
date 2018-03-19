package ssh_config

// UserKnownHostsFile
//         Specifies one or more files to use for the user host key data‐
//         base, separated by whitespace.  The default is
//         ~/.ssh/known_hosts, ~/.ssh/known_hosts2.
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

	sshInit(UserKnownHostsFile, funcInit)
	sshValid(UserKnownHostsFile, funcValid)
	sshParse(UserKnownHostsFile, funcParse)
}
