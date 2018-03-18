package ssh_config

// CanonicalizePermittedCNAMEs
//         Specifies rules to determine whether CNAMEs should be followed
//         when canonicalizing hostnames.  The rules consist of one or
//         more arguments of source_domain_list:target_domain_list, where
//         source_domain_list is a pattern-list of domains that may follow
//         CNAMEs in canonicalization, and target_domain_list is a pat‐
//         tern-list of domains that they may resolve to.
//
//         For example, “*.a.example.com:*.b.example.com,*.c.example.com”
//         will allow hostnames matching “*.a.example.com” to be canoni‐
//         calized to names in the “*.b.example.com” or “*.c.example.com”
//         domains.
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

	sshInit(CanonicalizePermittedCNAMEs, funcInit)
	sshValid(CanonicalizePermittedCNAMEs, funcValid)
	sshParse(CanonicalizePermittedCNAMEs, funcParse)
}
