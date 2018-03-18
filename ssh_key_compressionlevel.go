package ssh_config

// CompressionLevel
//         Specifies the compression level to use if compression is
//         enabled.  The argument must be an integer from 1 (fast) to 9
//         (slow, best).  The default level is 6, which is good for most
//         applications.  The meaning of the values is the same as in
//         gzip(1).  Note that this option applies to protocol version 1
//         only.
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

	sshInit(CompressionLevel, funcInit)
	sshValid(CompressionLevel, funcValid)
	sshParse(CompressionLevel, funcParse)
}
