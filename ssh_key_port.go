package ssh_config

import (
	"strconv"
	"strings"
)

// Port    Specifies the port number to connect on the remote host.  The
//         default is 22.
//

func init() {
	funcInit := func() (res string) {
		return "22"
	}

	funcValid := func(value string) (res bool) {
		value = strings.TrimSpace(value)
		port, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if port < 0 {
			return false
		}
		if port > 65535 {
			return false
		}
		return true
	}

	funcParse := func(input string) (values []string, err error) {
		// TODO
		panic("Not implemented")
		return
	}

	sshInit(Port, funcInit)
	sshValid(Port, funcValid)
	sshParse(Port, funcParse)
}
