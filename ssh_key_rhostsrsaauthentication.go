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
	ssh_init(RhostsRSAAuthentication, funcInit)
	ssh_valid(RhostsRSAAuthentication, funcValid)
}
