package ssh_config

// ProxyUseFdpass
//         Specifies that ProxyCommand will pass a connected file descrip‐
//         tor back to ssh(1) instead of continuing to execute and pass
//         data.  The default is “no”.
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

	sshInit(ProxyUseFdpass, funcInit)
	sshValid(ProxyUseFdpass, funcValid)
	sshParse(ProxyUseFdpass, funcParse)
}
