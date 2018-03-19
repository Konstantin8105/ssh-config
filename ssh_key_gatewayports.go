package ssh_config

// GatewayPorts
//         Specifies whether remote hosts are allowed to connect to local
//         forwarded ports.  By default, ssh(1) binds local port forward‐
//         ings to the loopback address.  This prevents other remote hosts
//         from connecting to forwarded ports.  GatewayPorts can be used
//         to specify that ssh should bind local port forwardings to the
//         wildcard address, thus allowing remote hosts to connect to for‐
//         warded ports.  The argument must be “yes” or “no”.  The default
//         is “no”.
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

	sshInit(GatewayPorts, funcInit)
	sshValid(GatewayPorts, funcValid)
	sshParse(GatewayPorts, funcParse)
}
