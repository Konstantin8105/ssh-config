package ssh_config

// IdentitiesOnly
//         Specifies that ssh(1) should only use the authentication iden‐
//         tity and certificate files explicitly configured in the
//         ssh_config files or passed on the ssh(1) command-line, even if
//         ssh-agent(1) or a PKCS11Provider offers more identities.  The
//         argument to this keyword must be “yes” or “no”.  This option is
//         intended for situations where ssh-agent offers many different
//         identities.  The default is “no”.
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
	ssh_init(IdentitiesOnly, funcInit)
	ssh_valid(IdentitiesOnly, funcValid)
}
