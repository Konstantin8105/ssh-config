package ssh_config

// SendEnv
//         Specifies what variables from the local environ(7) should be
//         sent to the server.  The server must also support it, and the
//         server must be configured to accept these environment vari‐
//         ables.  Note that the TERM environment variable is always sent
//         whenever a pseudo-terminal is requested as it is required by
//         the protocol.  Refer to AcceptEnv in sshd_config(5) for how to
//         configure the server.  Variables are specified by name, which
//         may contain wildcard characters.  Multiple environment vari‐
//         ables may be separated by whitespace or spread across multiple
//         SendEnv directives.  The default is not to send any environment
//         variables.
//
//         See PATTERNS for more information on patterns.
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

	sshInit(SendEnv, funcInit)
	sshValid(SendEnv, funcValid)
	sshParse(SendEnv, funcParse)
}
