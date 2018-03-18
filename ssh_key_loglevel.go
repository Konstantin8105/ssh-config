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
		return
	}
	funcValid := func() (res bool) {
		// TODO
		return
	}
	ssh_init(LogLevel, funcInit)
	ssh_valid(LogLevel, funcValid)
}
