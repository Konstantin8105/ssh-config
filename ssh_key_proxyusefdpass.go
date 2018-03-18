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
	ssh_init(ProxyUseFdpass, funcInit)
	ssh_valid(ProxyUseFdpass, funcValid)
}
