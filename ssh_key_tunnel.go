package ssh_config

// Tunnel  Request tun(4) device forwarding between the client and the
//         server.  The argument must be “yes”, “point-to-point” (layer
//         3), “ethernet” (layer 2), or “no”.  Specifying “yes” requests
//         the default tunnel mode, which is “point-to-point”.  The
//         default is “no”.
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

	sshInit(Tunnel, funcInit)
	sshValid(Tunnel, funcValid)
	sshParse(Tunnel, funcParse)
}
