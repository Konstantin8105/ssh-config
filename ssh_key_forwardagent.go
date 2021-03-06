package ssh_config

// ForwardAgent
//         Specifies whether the connection to the authentication agent
//         (if any) will be forwarded to the remote machine.  The argument
//         must be “yes” or “no”.  The default is “no”.
//
//         Agent forwarding should be enabled with caution.  Users with
//         the ability to bypass file permissions on the remote host (for
//         the agent's Unix-domain socket) can access the local agent
//         through the forwarded connection.  An attacker cannot obtain
//         key material from the agent, however they can perform opera‐
//         tions on the keys that enable them to authenticate using the
//         identities loaded into the agent.
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

	sshInit(ForwardAgent, funcInit)
	sshValid(ForwardAgent, funcValid)
	sshParse(ForwardAgent, funcParse)
}
