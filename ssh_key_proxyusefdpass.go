package ssh_config

// ProxyUseFdpass
//         Specifies that ProxyCommand will pass a connected file descrip‐
//         tor back to ssh(1) instead of continuing to execute and pass
//         data.  The default is “no”.
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
	ssh_init(ProxyUseFdpass, funcInit)
	ssh_valid(ProxyUseFdpass, funcValid)
}
