package ssh_config

// PKCS11Provider
//         Specifies which PKCS#11 provider to use.  The argument to this
//         keyword is the PKCS#11 shared library ssh(1) should use to com‐
//         municate with a PKCS#11 token providing the user's private RSA
//         key.
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
	ssh_init(PKCS11Provider, funcInit)
	ssh_valid(PKCS11Provider, funcValid)
}
