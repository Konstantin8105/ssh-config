package ssh_config

// PubkeyAcceptedKeyTypes
//         Specifies the key types that will be used for public key
//         authentication as a comma-separated pattern list.  Alternately
//         if the specified value begins with a ‘+’ character, then the
//         key types after it will be appended to the default instead of
//         replacing it.  The default for this option is:
//
//            ecdsa-sha2-nistp256-cert-v01@openssh.com,
//            ecdsa-sha2-nistp384-cert-v01@openssh.com,
//            ecdsa-sha2-nistp521-cert-v01@openssh.com,
//            ssh-ed25519-cert-v01@openssh.com,
//            ssh-rsa-cert-v01@openssh.com,
//            ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,
//            ssh-ed25519,ssh-rsa
//
//         The -Q option of ssh(1) may be used to list supported key
//         types.
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

	sshInit(PubkeyAcceptedKeyTypes, funcInit)
	sshValid(PubkeyAcceptedKeyTypes, funcValid)
	sshParse(PubkeyAcceptedKeyTypes, funcParse)
}
