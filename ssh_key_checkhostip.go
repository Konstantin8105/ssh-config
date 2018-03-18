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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(CheckHostIP, funcInit)
	ssh_valid(CheckHostIP, funcValid)
}
