package ssh_config

// ControlPersist
//         When used in conjunction with ControlMaster, specifies that the
//         master connection should remain open in the background (waiting
//         for future client connections) after the initial client connec‐
//         tion has been closed.  If set to “no”, then the master connec‐
//         tion will not be placed into the background, and will close as
//         soon as the initial client connection is closed.  If set to
//         “yes” or “0”, then the master connection will remain in the
//         background indefinitely (until killed or closed via a mechanism
//         such as the ssh(1) “-O exit” option).  If set to a time in sec‐
//         onds, or a time in any of the formats documented in
//         sshd_config(5), then the backgrounded master connection will
//         automatically terminate after it has remained idle (with no
//         client connections) for the specified time.
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
	ssh_init(ControlPersist, funcInit)
	ssh_valid(ControlPersist, funcValid)
}
