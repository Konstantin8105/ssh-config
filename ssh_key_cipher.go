package ssh_config

// Cipher  Specifies the cipher to use for encrypting the session in pro‐
//         tocol version 1.  Currently, “blowfish”, “3des”, and “des” are
//         supported.  des is only supported in the ssh(1) client for
//         interoperability with legacy protocol 1 implementations that do
//         not support the 3des cipher.  Its use is strongly discouraged
//         due to cryptographic weaknesses.  The default is “3des”.
//

func init() {
	funcInit := func() (res string) {
		// TODO
		panic("Not implemented")
		return
	}

	funcValid := func(value string) (res bool) {
		// TODO
		panic("Not implemented")
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		panic("Not implemented")
		return
	}

	sshInit(Cipher, funcInit)
	sshValid(Cipher, funcValid)
	sshParse(Cipher, funcParse)
}
