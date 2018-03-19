package ssh_config

// CheckHostIP
//         If this flag is set to “yes”, ssh(1) will additionally check
//         the host IP address in the known_hosts file.  This allows ssh
//         to detect if a host key changed due to DNS spoofing and will
//         add addresses of destination hosts to ~/.ssh/known_hosts in the
//         process, regardless of the setting of StrictHostKeyChecking.
//         If the option is set to “no”, the check will not be executed.
//         The default is “yes”.
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

	sshInit(CheckHostIP, funcInit)
	sshValid(CheckHostIP, funcValid)
	sshParse(CheckHostIP, funcParse)
}
