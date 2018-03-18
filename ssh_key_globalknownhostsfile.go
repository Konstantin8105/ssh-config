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
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(GlobalKnownHostsFile, funcInit)
	ssh_valid(GlobalKnownHostsFile, funcValid)
}
