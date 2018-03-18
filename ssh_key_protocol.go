package ssh_config

// Protocol
//         Specifies the protocol versions ssh(1) should support in order
//         of preference.  The possible values are ‘1’ and ‘2’.  Multiple
//         versions must be comma-separated.  When this option is set to
//         “2,1” ssh will try version 2 and fall back to version 1 if ver‐
//         sion 2 is not available.  The default is ‘2’.  Protocol 1 suf‐
//         fers from a number of cryptographic weaknesses and should not
//         be used.  It is only offered to support legacy devices.
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
	ssh_init(Protocol, funcInit)
	ssh_valid(Protocol, funcValid)
}
