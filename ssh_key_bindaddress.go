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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(BindAddress, funcInit)
	ssh_valid(BindAddress, funcValid)
}
