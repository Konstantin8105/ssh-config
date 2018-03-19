package ssh_config

var (
	mapInit  map[SSHKey]func() string
	mapValid map[SSHKey]func(string) bool
	mapParse map[SSHKey]func(string) ([]string, error)
)

func init() {
	mapInit = map[SSHKey]func() string{}
	mapValid = map[SSHKey]func(string) bool{}
	mapParse = map[SSHKey]func(string) ([]string, error){}
}

func sshInit(s SSHKey, f func() string) {
	mapInit[s] = f
}

func sshValid(s SSHKey, f func(string) bool) {
	mapValid[s] = f
}

func sshParse(s SSHKey, f func(string) ([]string, error)) {
	mapParse[s] = f
}
