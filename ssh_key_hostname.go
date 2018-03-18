package ssh_config

// HostName
//         Specifies the real host name to log into.  This can be used to
//         specify nicknames or abbreviations for hosts.  If the hostname
//         contains the character sequence ‘%h’, then this will be
//         replaced with the host name specified on the command line (this
//         is useful for manipulating unqualified names).  The character
//         sequence ‘%%’ will be replaced by a single ‘%’ character, which
//         may be used when specifying IPv6 link-local addresses.
//
//         The default is the name given on the command line.  Numeric IP
//         addresses are also permitted (both on the command line and in
//         HostName specifications).
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
	ssh_init(HostName, funcInit)
	ssh_valid(HostName, funcValid)
}
