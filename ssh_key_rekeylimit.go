package ssh_config

// RekeyLimit
//         Specifies the maximum amount of data that may be transmitted
//         before the session key is renegotiated, optionally followed a
//         maximum amount of time that may pass before the session key is
//         renegotiated.  The first argument is specified in bytes and may
//         have a suffix of ‘K’, ‘M’, or ‘G’ to indicate Kilobytes,
//         Megabytes, or Gigabytes, respectively.  The default is between
//         ‘1G’ and ‘4G’, depending on the cipher.  The optional second
//         value is specified in seconds and may use any of the units doc‐
//         umented in the TIME FORMATS section of sshd_config(5).  The
//         default value for RekeyLimit is “default none”, which means
//         that rekeying is performed after the cipher's default amount of
//         data has been sent or received and no time based rekeying is
//         done.
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

	sshInit(RekeyLimit, funcInit)
	sshValid(RekeyLimit, funcValid)
	sshParse(RekeyLimit, funcParse)
}
