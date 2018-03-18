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

	funcValid := func(value string) (res bool) {
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		return
	}

	sshInit(PKCS11Provider, funcInit)
	sshValid(PKCS11Provider, funcValid)
	sshParse(PKCS11Provider, funcParse)
}
