package ssh_config

// StreamLocalBindUnlink
//         Specifies whether to remove an existing Unix-domain socket file
//         for local or remote port forwarding before creating a new one.
//         If the socket file already exists and StreamLocalBindUnlink is
//         not enabled, ssh will be unable to forward the port to the
//         Unix-domain socket file.  This option is only used for port
//         forwarding to a Unix-domain socket file.
//
//         The argument must be “yes” or “no”.  The default is “no”.
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

	sshInit(StreamLocalBindUnlink, funcInit)
	sshValid(StreamLocalBindUnlink, funcValid)
	sshParse(StreamLocalBindUnlink, funcParse)
}
