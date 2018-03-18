package ssh_config

// RSAAuthentication
//         Specifies whether to try RSA authentication.  The argument to
//         this keyword must be “yes” or “no”.  RSA authentication will
//         only be attempted if the identity file exists, or an authenti‐
//         cation agent is running.  The default is “yes”.  Note that this
//         option applies to protocol version 1 only.
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

	sshInit(RSAAuthentication, funcInit)
	sshValid(RSAAuthentication, funcValid)
	sshParse(RSAAuthentication, funcParse)
}
