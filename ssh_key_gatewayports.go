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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(GatewayPorts, funcInit)
	ssh_valid(GatewayPorts, funcValid)
}
