package ssh_config

// XAuthLocation
//         Specifies the full pathname of the xauth(1) program.  The
//         default is /usr/bin/xauth.
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

	sshInit(XAuthLocation, funcInit)
	sshValid(XAuthLocation, funcValid)
	sshParse(XAuthLocation, funcParse)
}
