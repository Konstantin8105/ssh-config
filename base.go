package ssh_config

import (
	"fmt"
	"strings"
)

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
	// checking of init-valid
	// init-value must be valid
	initFunc, ok := mapInit[s]
	if !ok {
		panic(fmt.Errorf("Before initialization of valid-function need to "+
			"initialize init-function for ssh-key: %v", s))
	}
	if !f(initFunc()) {
		panic(fmt.Errorf("Not valid init value for ssh-key : %v", s))
	}

	// add valid-function to map
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

func parseSingleString(input string) (values []string, err error) {
	input = strings.TrimSpace(input)
	if strings.Count(input, " ") > 0 {
		err = fmt.Errorf("Not acceptable more then 1 value ` `: `%v`",
			input)
	}
	if strings.Count(input, ",") > 0 {
		err = fmt.Errorf("Not acceptable more then 1 value `,`: `%v`",
			input)
	}
	values = append(values, input)
	return
}
