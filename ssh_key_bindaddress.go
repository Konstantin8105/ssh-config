package ssh_config

// BindAddress
//         Use the specified address on the local machine as the source
//         address of the connection.  Only useful on systems with more
//         than one address.  Note that this option does not work if
//         UsePrivilegedPort is set to “yes”.
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

	sshInit(BindAddress, funcInit)
	sshValid(BindAddress, funcValid)
	sshParse(BindAddress, funcParse)
}
