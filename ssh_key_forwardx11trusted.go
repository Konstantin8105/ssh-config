package ssh_config

func init() {
	funcInit := func() (res string) {
		// TODO
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(ForwardX11Trusted, funcInit)
	ssh_valid(ForwardX11Trusted, funcValid)
}
