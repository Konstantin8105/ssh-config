package ssh_config

// ExitOnForwardFailure
//         Specifies whether ssh(1) should terminate the connection if it
//         cannot set up all requested dynamic, tunnel, local, and remote
//         port forwardings, (e.g. if either end is unable to bind and
//         listen on a specified port).  Note that ExitOnForwardFailure
//         does not apply to connections made over port forwardings and
//         will not, for example, cause ssh(1) to exit if TCP connections
//         to the ultimate forwarding destination fail.  The argument must
//         be “yes” or “no”.  The default is “no”.
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

	sshInit(ExitOnForwardFailure, funcInit)
	sshValid(ExitOnForwardFailure, funcValid)
	sshParse(ExitOnForwardFailure, funcParse)
}
