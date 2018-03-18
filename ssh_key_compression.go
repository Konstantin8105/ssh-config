package ssh_config

// Compression
//         Specifies whether to use compression.  The argument must be
//         “yes” or “no”.  The default is “no”.
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
	ssh_init(Compression, funcInit)
	ssh_valid(Compression, funcValid)
}
