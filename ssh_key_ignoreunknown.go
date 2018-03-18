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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(IgnoreUnknown, funcInit)
	ssh_valid(IgnoreUnknown, funcValid)
}
