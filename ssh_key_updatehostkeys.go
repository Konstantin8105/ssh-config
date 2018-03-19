package ssh_config

// UpdateHostKeys
//         Specifies whether ssh(1) should accept notifications of addi‐
//         tional hostkeys from the server sent after authentication has
//         completed and add them to UserKnownHostsFile.  The argument
//         must be “yes”, “no” (the default) or “ask”.  Enabling this
//         option allows learning alternate hostkeys for a server and sup‐
//         ports graceful key rotation by allowing a server to send
//         replacement public keys before old ones are removed.  Addi‐
//         tional hostkeys are only accepted if the key used to authenti‐
//         cate the host was already trusted or explicitly accepted by the
//         user.  If UpdateHostKeys is set to “ask”, then the user is
//         asked to confirm the modifications to the known_hosts file.
//         Confirmation is currently incompatible with ControlPersist, and
//         will be disabled if it is enabled.
//
//         Presently, only sshd(8) from OpenSSH 6.8 and greater support
//         the “hostkeys@openssh.com” protocol extension used to inform
//         the client of all the server's hostkeys.
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

	sshInit(UpdateHostKeys, funcInit)
	sshValid(UpdateHostKeys, funcValid)
	sshParse(UpdateHostKeys, funcParse)
}
