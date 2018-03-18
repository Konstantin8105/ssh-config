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
	ssh_init(ForwardX11, funcInit)
	ssh_valid(ForwardX11, funcValid)
}
