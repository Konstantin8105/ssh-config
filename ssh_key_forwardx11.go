package ssh_config

// ForwardX11
//         Specifies whether X11 connections will be automatically redi‐
//         rected over the secure channel and DISPLAY set.  The argument
//         must be “yes” or “no”.  The default is “no”.
//
//         X11 forwarding should be enabled with caution.  Users with the
//         ability to bypass file permissions on the remote host (for the
//         user's X11 authorization database) can access the local X11
//         display through the forwarded connection.  An attacker may then
//         be able to perform activities such as keystroke monitoring if
//         the ForwardX11Trusted option is also enabled.
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

	sshInit(ForwardX11, funcInit)
	sshValid(ForwardX11, funcValid)
	sshParse(ForwardX11, funcParse)
}
