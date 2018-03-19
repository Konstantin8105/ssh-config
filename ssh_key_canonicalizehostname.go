package ssh_config

// CanonicalizeHostname
//         Controls whether explicit hostname canonicalization is per‐
//         formed.  The default, “no”, is not to perform any name rewrit‐
//         ing and let the system resolver handle all hostname lookups.
//         If set to “yes” then, for connections that do not use a
//         ProxyCommand, ssh(1) will attempt to canonicalize the hostname
//         specified on the command line using the CanonicalDomains suf‐
//         fixes and CanonicalizePermittedCNAMEs rules.  If
//         CanonicalizeHostname is set to “always”, then canonicalization
//         is applied to proxied connections too.
//
//         If this option is enabled, then the configuration files are
//         processed again using the new target name to pick up any new
//         configuration in matching Host and Match stanzas.
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

	sshInit(CanonicalizeHostname, funcInit)
	sshValid(CanonicalizeHostname, funcValid)
	sshParse(CanonicalizeHostname, funcParse)
}
