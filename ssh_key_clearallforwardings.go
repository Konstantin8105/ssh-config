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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(ClearAllForwardings, funcInit)
	ssh_valid(ClearAllForwardings, funcValid)
}
