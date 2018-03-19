package ssh_config

// MACs    Specifies the MAC (message authentication code) algorithms in
//         order of preference.  The MAC algorithm is used for data
//         integrity protection.  Multiple algorithms must be comma-sepa‐
//         rated.  If the specified value begins with a ‘+’ character,
//         then the specified algorithms will be appended to the default
//         set instead of replacing them.
//
//         The algorithms that contain “-etm” calculate the MAC after
//         encryption (encrypt-then-mac).  These are considered safer and
//         their use recommended.
//
//         The default is:
//
//               umac-64-etm@openssh.com,umac-128-etm@openssh.com,
//               hmac-sha2-256-etm@openssh.com,hmac-sha2-512-etm@openssh.com,
//               hmac-sha1-etm@openssh.com,
//               umac-64@openssh.com,umac-128@openssh.com,
//               hmac-sha2-256,hmac-sha2-512,hmac-sha1
//
//         The list of available MAC algorithms may also be obtained using
//         the -Q option of ssh(1) with an argument of “mac”.
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

	sshInit(MACs, funcInit)
	sshValid(MACs, funcValid)
	sshParse(MACs, funcParse)
}
