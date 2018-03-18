package ssh_config

// UserKnownHostsFile
//         Specifies one or more files to use for the user host key data‐
//         base, separated by whitespace.  The default is
//         ~/.ssh/known_hosts, ~/.ssh/known_hosts2.
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
	ssh_init(UserKnownHostsFile, funcInit)
	ssh_valid(UserKnownHostsFile, funcValid)
}
