package ssh_config

// KexAlgorithms
//         Specifies the available KEX (Key Exchange) algorithms.  Multi‐
//         ple algorithms must be comma-separated.  Alternately if the
//         specified value begins with a ‘+’ character, then the specified
//         methods will be appended to the default set instead of replac‐
//         ing them.  The default is:
//
//               curve25519-sha256@libssh.org,
//               ecdh-sha2-nistp256,ecdh-sha2-nistp384,ecdh-sha2-nistp521,
//               diffie-hellman-group-exchange-sha256,
//               diffie-hellman-group-exchange-sha1,
//               diffie-hellman-group14-sha1
//
//         The list of available key exchange algorithms may also be
//         obtained using the -Q option of ssh(1) with an argument of
//         “kex”.
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

	sshInit(KexAlgorithms, funcInit)
	sshValid(KexAlgorithms, funcValid)
	sshParse(KexAlgorithms, funcParse)
}
