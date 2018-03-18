package ssh_config

// TunnelDevice
//         Specifies the tun(4) devices to open on the client (local_tun)
//         and the server (remote_tun).
//
//         The argument must be local_tun[:remote_tun].  The devices may
//         be specified by numerical ID or the keyword “any”, which uses
//         the next available tunnel device.  If remote_tun is not speci‐
//         fied, it defaults to “any”.  The default is “any:any”.
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
	ssh_init(TunnelDevice, funcInit)
	ssh_valid(TunnelDevice, funcValid)
}
