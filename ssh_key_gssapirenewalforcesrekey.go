package ssh_config

// GSSAPIRenewalForcesRekey
//         If set to “yes” then renewal of the client's GSSAPI credentials
//         will force the rekeying of the ssh connection. With a compati‐
//         ble server, this can delegate the renewed credentials to a ses‐
//         sion on the server.  The default is “no”.
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

	sshInit(GSSAPIRenewalForcesRekey, funcInit)
	sshValid(GSSAPIRenewalForcesRekey, funcValid)
	sshParse(GSSAPIRenewalForcesRekey, funcParse)
}
