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
	ssh_init(AddKeysToAgent, funcInit)
	ssh_valid(AddKeysToAgent, funcValid)
}
