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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(StreamLocalBindUnlink, funcInit)
	ssh_valid(StreamLocalBindUnlink, funcValid)
}
