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
	ssh_init(PKCS11Provider, funcInit)
	ssh_valid(PKCS11Provider, funcValid)
}
