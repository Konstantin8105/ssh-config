package ssh_config

var (
	mapInit  map[SSHKey]func() string
	mapValid map[SSHKey]func() bool
)

func ssh_init(s SSHKey, f func() string) {
	mapInit[s] = f
}

func ssh_valid(s SSHKey, f func() bool) {
	mapValid[s] = f
}
