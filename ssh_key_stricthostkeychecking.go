package ssh_config

// StrictHostKeyChecking
//         If this flag is set to “yes”, ssh(1) will never automatically
//         add host keys to the ~/.ssh/known_hosts file, and refuses to
//         connect to hosts whose host key has changed.  This provides
//         maximum protection against trojan horse attacks, though it can
//         be annoying when the /etc/ssh/ssh_known_hosts file is poorly
//         maintained or when connections to new hosts are frequently
//         made.  This option forces the user to manually add all new
//         hosts.  If this flag is set to “no”, ssh will automatically add
//         new host keys to the user known hosts files.  If this flag is
//         set to “ask”, new host keys will be added to the user known
//         host files only after the user has confirmed that is what they
//         really want to do, and ssh will refuse to connect to hosts
//         whose host key has changed.  The host keys of known hosts will
//         be verified automatically in all cases.  The argument must be
//         “yes”, “no”, or “ask”.  The default is “ask”.
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

	sshInit(StrictHostKeyChecking, funcInit)
	sshValid(StrictHostKeyChecking, funcValid)
	sshParse(StrictHostKeyChecking, funcParse)
}
