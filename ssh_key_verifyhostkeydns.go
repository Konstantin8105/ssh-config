package ssh_config

// VerifyHostKeyDNS
//         Specifies whether to verify the remote key using DNS and SSHFP
//         resource records.  If this option is set to “yes”, the client
//         will implicitly trust keys that match a secure fingerprint from
//         DNS.  Insecure fingerprints will be handled as if this option
//         was set to “ask”.  If this option is set to “ask”, information
//         on fingerprint match will be displayed, but the user will still
//         need to confirm new host keys according to the
//         StrictHostKeyChecking option.  The argument must be “yes”,
//         “no”, or “ask”.  The default is “no”.
//
//         See also VERIFYING HOST KEYS in ssh(1).
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

	sshInit(VerifyHostKeyDNS, funcInit)
	sshValid(VerifyHostKeyDNS, funcValid)
	sshParse(VerifyHostKeyDNS, funcParse)
}
