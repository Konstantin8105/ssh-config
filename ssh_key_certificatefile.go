package ssh_config

// CertificateFile
//         Specifies a file from which the user's certificate is read.  A
//         corresponding private key must be provided separately in order
//         to use this certificate either from an IdentityFile directive
//         or -i flag to ssh(1), via ssh-agent(1), or via a
//         PKCS11Provider.
//
//         The file name may use the tilde syntax to refer to a user's
//         home directory or one of the following escape characters: ‘%d’
//         (local user's home directory), ‘%u’ (local user name), ‘%l’
//         (local host name), ‘%h’ (remote host name) or ‘%r’ (remote user
//         name).
//
//         It is possible to have multiple certificate files specified in
//         configuration files; these certificates will be tried in
//         sequence.  Multiple CertificateFile directives will add to the
//         list of certificates used for authentication.
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
	ssh_init(CertificateFile, funcInit)
	ssh_valid(CertificateFile, funcValid)
}
