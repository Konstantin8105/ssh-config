package ssh_config

// FingerprintHash
//         Specifies the hash algorithm used when displaying key finger‐
//         prints.  Valid options are: “md5” and “sha256”.  The default is
//         “sha256”.
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
	ssh_init(FingerprintHash, funcInit)
	ssh_valid(FingerprintHash, funcValid)
}
