package ssh_config

// PermitLocalCommand
//         Allow local command execution via the LocalCommand option or
//         using the !command escape sequence in ssh(1).  The argument
//         must be “yes” or “no”.  The default is “no”.
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

	sshInit(PermitLocalCommand, funcInit)
	sshValid(PermitLocalCommand, funcValid)
	sshParse(PermitLocalCommand, funcParse)
}
