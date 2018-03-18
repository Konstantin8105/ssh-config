package ssh_config

// XAuthLocation
//         Specifies the full pathname of the xauth(1) program.  The
//         default is /usr/bin/xauth.
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
	ssh_init(XAuthLocation, funcInit)
	ssh_valid(XAuthLocation, funcValid)
}
