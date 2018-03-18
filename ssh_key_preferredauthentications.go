package ssh_config

// PreferredAuthentications
//         Specifies the order in which the client should try authentica‐
//         tion methods.  This allows a client to prefer one method (e.g.
//         keyboard-interactive) over another method (e.g. password).  The
//         default is:
//
//               gssapi-with-mic,hostbased,publickey,
//               keyboard-interactive,password
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
	ssh_init(PreferredAuthentications, funcInit)
	ssh_valid(PreferredAuthentications, funcValid)
}
