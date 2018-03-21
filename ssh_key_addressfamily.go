package ssh_config

// AddressFamily
//         Specifies which address family to use when connecting.  Valid
//         arguments are “any”, “inet” (use IPv4 only), or “inet6” (use
//         IPv6 only).  The default is “any”.
//

func init() {
	funcInit := func() (res string) {
		return "any"
	}

	funcValid := func(value string) (res bool) {
		return isValid(value, "any", "inet", "inet6")
	}

	funcParse := func(input string) (values []string, err error) {
		return parseSingleString(input)
	}

	sshInit(AddressFamily, funcInit)
	sshValid(AddressFamily, funcValid)
	sshParse(AddressFamily, funcParse)
}
