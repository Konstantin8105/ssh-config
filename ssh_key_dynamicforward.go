package ssh_config

// DynamicForward
//         Specifies that a TCP port on the local machine be forwarded
//         over the secure channel, and the application protocol is then
//         used to determine where to connect to from the remote machine.
//
//         The argument must be [bind_address:]port.  IPv6 addresses can
//         be specified by enclosing addresses in square brackets.  By
//         default, the local port is bound in accordance with the
//         GatewayPorts setting.  However, an explicit bind_address may be
//         used to bind the connection to a specific address.  The
//         bind_address of “localhost” indicates that the listening port
//         be bound for local use only, while an empty address or ‘*’
//         indicates that the port should be available from all inter‐
//         faces.
//
//         Currently the SOCKS4 and SOCKS5 protocols are supported, and
//         ssh(1) will act as a SOCKS server.  Multiple forwardings may be
//         specified, and additional forwardings can be given on the com‐
//         mand line.  Only the superuser can forward privileged ports.
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
	ssh_init(DynamicForward, funcInit)
	ssh_valid(DynamicForward, funcValid)
}
