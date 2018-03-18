package ssh_config

// HashKnownHosts
//         Indicates that ssh(1) should hash host names and addresses when
//         they are added to ~/.ssh/known_hosts.  These hashed names may
//         be used normally by ssh(1) and sshd(8), but they do not reveal
//         identifying information should the file's contents be dis‐
//         closed.  The default is “no”.  Note that existing names and
//         addresses in known hosts files will not be converted automati‐
//         cally, but may be manually hashed using ssh-keygen(1).  Use of
//         this option may break facilities such as tab-completion that
//         rely on being able to read unhashed host names from
//         ~/.ssh/known_hosts.
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
	ssh_init(HashKnownHosts, funcInit)
	ssh_valid(HashKnownHosts, funcValid)
}
