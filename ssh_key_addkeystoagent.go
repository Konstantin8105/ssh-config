package ssh_config

// AddKeysToAgent
//         Specifies whether keys should be automatically added to a run‐
//         ning ssh-agent(1).  If this option is set to “yes” and a key is
//         loaded from a file, the key and its passphrase are added to the
//         agent with the default lifetime, as if by ssh-add(1).  If this
//         option is set to “ask”, ssh will require confirmation using the
//         SSH_ASKPASS program before adding a key (see ssh-add(1) for
//         details).  If this option is set to “confirm”, each use of the
//         key must be confirmed, as if the -c option was specified to
//         ssh-add(1).  If this option is set to “no”, no keys are added
//         to the agent.  The argument must be “yes”, “confirm”, “ask”, or
//         “no”.  The default is “no”.
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

	sshInit(AddKeysToAgent, funcInit)
	sshValid(AddKeysToAgent, funcValid)
	sshParse(AddKeysToAgent, funcParse)
}
