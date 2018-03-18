package ssh_config

// NumberOfPasswordPrompts
//         Specifies the number of password prompts before giving up.  The
//         argument to this keyword must be an integer.  The default is 3.
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
	ssh_init(NumberOfPasswordPrompts, funcInit)
	ssh_valid(NumberOfPasswordPrompts, funcValid)
}
