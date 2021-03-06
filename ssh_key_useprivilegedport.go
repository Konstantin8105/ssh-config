package ssh_config

// UsePrivilegedPort
//         Specifies whether to use a privileged port for outgoing connec‐
//         tions.  The argument must be “yes” or “no”.  The default is
//         “no”.  If set to “yes”, ssh(1) must be setuid root.  Note that
//         this option must be set to “yes” for RhostsRSAAuthentication
//         with older servers.
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

	sshInit(UsePrivilegedPort, funcInit)
	sshValid(UsePrivilegedPort, funcValid)
	sshParse(UsePrivilegedPort, funcParse)
}
