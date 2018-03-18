package ssh_config

// Ciphers
//         Specifies the ciphers allowed for protocol version 2 in order
//         of preference.  Multiple ciphers must be comma-separated.  If
//         the specified value begins with a ‘+’ character, then the spec‐
//         ified ciphers will be appended to the default set instead of
//         replacing them.
//
//         The supported ciphers are:
//
//               3des-cbc
//               aes128-cbc
//               aes192-cbc
//               aes256-cbc
//               aes128-ctr
//               aes192-ctr
//               aes256-ctr
//               aes128-gcm@openssh.com
//               aes256-gcm@openssh.com
//               arcfour
//               arcfour128
//               arcfour256
//               blowfish-cbc
//               cast128-cbc
//               chacha20-poly1305@openssh.com
//
//         The default is:
//
//               chacha20-poly1305@openssh.com,
//               aes128-ctr,aes192-ctr,aes256-ctr,
//               aes128-gcm@openssh.com,aes256-gcm@openssh.com,
//               aes128-cbc,aes192-cbc,aes256-cbc,3des-cbc
//
//         The list of available ciphers may also be obtained using the -Q
//         option of ssh(1) with an argument of “cipher”.
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

	sshInit(Ciphers, funcInit)
	sshValid(Ciphers, funcValid)
	sshParse(Ciphers, funcParse)
}
