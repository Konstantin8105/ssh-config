package ssh_config

// IgnoreUnknown
//         Specifies a pattern-list of unknown options to be ignored if
//         they are encountered in configuration parsing.  This may be
//         used to suppress errors if ssh_config contains options that are
//         unrecognised by ssh(1).  It is recommended that IgnoreUnknown
//         be listed early in the configuration file as it will not be
//         applied to unknown options that appear before it.
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

	sshInit(IgnoreUnknown, funcInit)
	sshValid(IgnoreUnknown, funcValid)
	sshParse(IgnoreUnknown, funcParse)
}
