package ssh_config

// User    Specifies the user to log in as.  This can be useful when a
//         different user name is used on different machines.  This saves
//         the trouble of having to remember to give the user name on the
//         command line.
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
	ssh_init(User, funcInit)
	ssh_valid(User, funcValid)
}
