package ssh_config

// ControlMaster
//         Enables the sharing of multiple sessions over a single network
//         connection.  When set to “yes”, ssh(1) will listen for connec‐
//         tions on a control socket specified using the ControlPath argu‐
//         ment.  Additional sessions can connect to this socket using the
//         same ControlPath with ControlMaster set to “no” (the default).
//         These sessions will try to reuse the master instance's network
//         connection rather than initiating new ones, but will fall back
//         to connecting normally if the control socket does not exist, or
//         is not listening.
//
//         Setting this to “ask” will cause ssh to listen for control con‐
//         nections, but require confirmation using ssh-askpass(1).  If
//         the ControlPath cannot be opened, ssh will continue without
//         connecting to a master instance.
//
//         X11 and ssh-agent(1) forwarding is supported over these multi‐
//         plexed connections, however the display and agent forwarded
//         will be the one belonging to the master connection i.e. it is
//         not possible to forward multiple displays or agents.
//
//         Two additional options allow for opportunistic multiplexing:
//         try to use a master connection but fall back to creating a new
//         one if one does not already exist.  These options are: “auto”
//         and “autoask”.  The latter requires confirmation like the “ask”
//         option.
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

	sshInit(ControlMaster, funcInit)
	sshValid(ControlMaster, funcValid)
	sshParse(ControlMaster, funcParse)
}
