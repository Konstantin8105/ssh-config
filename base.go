package ssh_config

import "strings"

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

func isValid(input string, allowableValues ...string) bool {
	if len(allowableValues) == 0 {
		panic("Values is empty")
	}
	input = strings.TrimSpace(input)
	for _, av := range allowableValues {
		if input == strings.TrimSpace(av) {
			return true
		}
	}
	return false
}
