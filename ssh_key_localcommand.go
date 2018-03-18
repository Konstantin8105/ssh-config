package ssh_config

// LocalCommand
//         Specifies a command to execute on the local machine after suc‐
//         cessfully connecting to the server.  The command string extends
//         to the end of the line, and is executed with the user's shell.
//         The following escape character substitutions will be performed:
//         ‘%d’ (local user's home directory), ‘%h’ (remote host name),
//         ‘%l’ (local host name), ‘%n’ (host name as provided on the com‐
//         mand line), ‘%p’ (remote port), ‘%r’ (remote user name) or ‘%u’
//         (local user name) or ‘%C’ by a hash of the concatenation:
//         %l%h%p%r.
//
//         The command is run synchronously and does not have access to
//         the session of the ssh(1) that spawned it.  It should not be
//         used for interactive commands.
//
//         This directive is ignored unless PermitLocalCommand has been
//         enabled.
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
	ssh_init(LocalCommand, funcInit)
	ssh_valid(LocalCommand, funcValid)
}
