package ssh_config

// ForwardX11Timeout
//         Specify a timeout for untrusted X11 forwarding using the format
//         described in the TIME FORMATS section of sshd_config(5).  X11
//         connections received by ssh(1) after this time will be refused.
//         The default is to disable untrusted X11 forwarding after twenty
//         minutes has elapsed.
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
	ssh_init(ForwardX11Timeout, funcInit)
	ssh_valid(ForwardX11Timeout, funcValid)
}
