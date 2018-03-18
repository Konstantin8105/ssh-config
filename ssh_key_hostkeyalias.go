package ssh_config

// HostKeyAlias
//         Specifies an alias that should be used instead of the real host
//         name when looking up or saving the host key in the host key
//         database files.  This option is useful for tunneling SSH con‐
//         nections or for multiple servers running on a single host.
//

func init() {
	funcInit := func() (res string) {
		// TODO
		return
	}

	funcValid := func(value string) (res bool) {
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		return
	}

	sshInit(HostKeyAlias, funcInit)
	sshValid(HostKeyAlias, funcValid)
	sshParse(HostKeyAlias, funcParse)
}
