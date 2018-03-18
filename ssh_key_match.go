package ssh_config

// Match   Restricts the following declarations (up to the next Host or
//         Match keyword) to be used only when the conditions following
//         the Match keyword are satisfied.  Match conditions are speci‐
//         fied using one or more criteria or the single token all which
//         always matches.  The available criteria keywords are:
//         canonical, exec, host, originalhost, user, and localuser.  The
//         all criteria must appear alone or immediately after canonical.
//         Other criteria may be combined arbitrarily.  All criteria but
//         all and canonical require an argument.  Criteria may be negated
//         by prepending an exclamation mark (‘!’).
//
//         The canonical keyword matches only when the configuration file
//         is being re-parsed after hostname canonicalization (see the
//         CanonicalizeHostname option.)  This may be useful to specify
//         conditions that work with canonical host names only.  The exec
//         keyword executes the specified command under the user's shell.
//         If the command returns a zero exit status then the condition is
//         considered true.  Commands containing whitespace characters
//         must be quoted.  The following character sequences in the com‐
//         mand will be expanded prior to execution: ‘%L’ will be substi‐
//         tuted by the first component of the local host name, ‘%l’ will
//         be substituted by the local host name (including any domain
//         name), ‘%h’ will be substituted by the target host name, ‘%n’
//         will be substituted by the original target host name specified
//         on the command-line, ‘%p’ the destination port, ‘%r’ by the
//         remote login username, and ‘%u’ by the username of the user
//         running ssh(1).
//
//         The other keywords' criteria must be single entries or comma-
//         separated lists and may use the wildcard and negation operators
//         described in the PATTERNS section.  The criteria for the host
//         keyword are matched against the target hostname, after any sub‐
//         stitution by the Hostname or CanonicalizeHostname options.  The
//         originalhost keyword matches against the hostname as it was
//         specified on the command-line.  The user keyword matches
//         against the target username on the remote host.  The localuser
//         keyword matches against the name of the local user running
//         ssh(1) (this keyword may be useful in system-wide ssh_config
//         files).
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
	ssh_init(Match, funcInit)
	ssh_valid(Match, funcValid)
}
