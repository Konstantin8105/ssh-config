package ssh_config

// ForwardX11Trusted
//         If this option is set to “yes”, remote X11 clients will have
//         full access to the original X11 display.
//
//         If this option is set to “no”, remote X11 clients will be con‐
//         sidered untrusted and prevented from stealing or tampering with
//         data belonging to trusted X11 clients.  Furthermore, the
//         xauth(1) token used for the session will be set to expire after
//         20 minutes.  Remote clients will be refused access after this
//         time.
//
//         The default is “yes” (Debian-specific).
//
//         See the X11 SECURITY extension specification for full details
//         on the restrictions imposed on untrusted clients.
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

	sshInit(ForwardX11Trusted, funcInit)
	sshValid(ForwardX11Trusted, funcValid)
	sshParse(ForwardX11Trusted, funcParse)
}
