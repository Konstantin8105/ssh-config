package ssh_config

// LogLevel
//         Gives the verbosity level that is used when logging messages
//         from ssh(1).  The possible values are: QUIET, FATAL, ERROR,
//         INFO, VERBOSE, DEBUG, DEBUG1, DEBUG2, and DEBUG3.  The default
//         is INFO.  DEBUG and DEBUG1 are equivalent.  DEBUG2 and DEBUG3
//         each specify higher levels of verbose output.
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

	sshInit(LogLevel, funcInit)
	sshValid(LogLevel, funcValid)
	sshParse(LogLevel, funcParse)
}
