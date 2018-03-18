package ssh_config

// RevokedHostKeys
//         Specifies revoked host public keys.  Keys listed in this file
//         will be refused for host authentication.  Note that if this
//         file does not exist or is not readable, then host authentica‐
//         tion will be refused for all hosts.  Keys may be specified as a
//         text file, listing one public key per line, or as an OpenSSH
//         Key Revocation List (KRL) as generated by ssh-keygen(1).  For
//         more information on KRLs, see the KEY REVOCATION LISTS section
//         in ssh-keygen(1).
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
	ssh_init(RevokedHostKeys, funcInit)
	ssh_valid(RevokedHostKeys, funcValid)
}
