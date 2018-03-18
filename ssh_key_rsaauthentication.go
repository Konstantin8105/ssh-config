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
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(RSAAuthentication, funcInit)
	ssh_valid(RSAAuthentication, funcValid)
}
