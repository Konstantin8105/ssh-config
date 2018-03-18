package ssh_config

// HostKeyAlgorithms
//         Specifies the host key algorithms that the client wants to use
//         in order of preference.  Alternately if the specified value
//         begins with a ‘+’ character, then the specified key types will
//         be appended to the default set instead of replacing them.  The
//         default for this option is:
//
//            ecdsa-sha2-nistp256-cert-v01@openssh.com,
//            ecdsa-sha2-nistp384-cert-v01@openssh.com,
//            ecdsa-sha2-nistp521-cert-v01@openssh.com,
//            ssh-ed25519-cert-v01@openssh.com,
//            ssh-rsa-cert-v01@openssh.com,
//            ecdsa-sha2-nistp256,ecdsa-sha2-nistp384,ecdsa-sha2-nistp521,
//            ssh-ed25519,ssh-rsa
//
//         If hostkeys are known for the destination host then this
//         default is modified to prefer their algorithms.
//
//         The list of available key types may also be obtained using the
//         -Q option of ssh(1) with an argument of “key”.
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
	ssh_init(HostKeyAlgorithms, funcInit)
	ssh_valid(HostKeyAlgorithms, funcValid)
}
