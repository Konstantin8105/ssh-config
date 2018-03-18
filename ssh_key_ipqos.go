package ssh_config

// IPQoS   Specifies the IPv4 type-of-service or DSCP class for connec‐
//         tions.  Accepted values are “af11”, “af12”, “af13”, “af21”,
//         “af22”, “af23”, “af31”, “af32”, “af33”, “af41”, “af42”, “af43”,
//         “cs0”, “cs1”, “cs2”, “cs3”, “cs4”, “cs5”, “cs6”, “cs7”, “ef”,
//         “lowdelay”, “throughput”, “reliability”, or a numeric value.
//         This option may take one or two arguments, separated by white‐
//         space.  If one argument is specified, it is used as the packet
//         class unconditionally.  If two values are specified, the first
//         is automatically selected for interactive sessions and the sec‐
//         ond for non-interactive sessions.  The default is “lowdelay”
//         for interactive sessions and “throughput” for non-interactive
//         sessions.
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

	sshInit(IPQoS, funcInit)
	sshValid(IPQoS, funcValid)
	sshParse(IPQoS, funcParse)
}
