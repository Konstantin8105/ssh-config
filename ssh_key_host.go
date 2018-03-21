package ssh_config

import "strings"

// Host    Restricts the following declarations (up to the next Host or
//         Match keyword) to be only for those hosts that match one of the
//         patterns given after the keyword.  If more than one pattern is
//         provided, they should be separated by whitespace.  A single ‘*’
//         as a pattern can be used to provide global defaults for all
//         hosts.  The host is usually the hostname argument given on the
//         command line (see the CanonicalizeHostname option for excep‐
//         tions.)
//
//         A pattern entry may be negated by prefixing it with an exclama‐
//         tion mark (‘!’).  If a negated entry is matched, then the Host
//         entry is ignored, regardless of whether any other patterns on
//         the line match.  Negated matches are therefore useful to pro‐
//         vide exceptions for wildcard matches.
//
//         See PATTERNS for more information on patterns.
//

func init() {
	funcInit := func() (res string) {
		// TODO
		panic("Not implemented")
		return
	}

	funcValid := func(value string) (res bool) {
		if value == "*" {
			return true
		}

		// TODO
		// panic("Not implemented")
		return
	}

	funcParse := func(input string) (values []string, err error) {
		values = strings.Split(input, " ")
		for i := range values {
			values[i] = strings.TrimSpace(values[i])
		}
		return
	}

	sshInit(Host, funcInit)
	sshValid(Host, funcValid)
	sshParse(Host, funcParse)
}
