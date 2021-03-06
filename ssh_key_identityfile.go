package ssh_config

// IdentityFile
//         Specifies a file from which the user's DSA, ECDSA, Ed25519 or
//         RSA authentication identity is read.  The default is
//         ~/.ssh/identity for protocol version 1, and ~/.ssh/id_dsa,
//         ~/.ssh/id_ecdsa, ~/.ssh/id_ed25519 and ~/.ssh/id_rsa for proto‐
//         col version 2.  Additionally, any identities represented by the
//         authentication agent will be used for authentication unless
//         IdentitiesOnly is set.  If no certificates have been explicitly
//         specified by CertificateFile, ssh(1) will try to load certifi‐
//         cate information from the filename obtained by appending
//         -cert.pub to the path of a specified IdentityFile.
//
//         The file name may use the tilde syntax to refer to a user's
//         home directory or one of the following escape characters: ‘%d’
//         (local user's home directory), ‘%u’ (local user name), ‘%l’
//         (local host name), ‘%h’ (remote host name) or ‘%r’ (remote user
//         name).
//
//         It is possible to have multiple identity files specified in
//         configuration files; all these identities will be tried in
//         sequence.  Multiple IdentityFile directives will add to the
//         list of identities tried (this behaviour differs from that of
//         other configuration directives).
//
//         IdentityFile may be used in conjunction with IdentitiesOnly to
//         select which identities in an agent are offered during authen‐
//         tication.  IdentityFile may also be used in conjunction with
//         CertificateFile in order to provide any certificate also needed
//         for authentication with the identity.
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

	sshInit(IdentityFile, funcInit)
	sshValid(IdentityFile, funcValid)
	sshParse(IdentityFile, funcParse)
}
