package ssh_config

// Port    Specifies the port number to connect on the remote host.  The
//         default is 22.
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
	ssh_init(Port, funcInit)
	ssh_valid(Port, funcValid)
}
