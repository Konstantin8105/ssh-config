package ssh_config

import (
	"fmt"
	"strconv"
	"strings"
)

// ServerAliveInterval
//         Sets a timeout interval in seconds after which if no data has
//         been received from the server, ssh(1) will send a message
//         through the encrypted channel to request a response from the
//         server.  The default is 0, indicating that these messages will
//         not be sent to the server, or 300 if the BatchMode option is
//         set.  ProtocolKeepAlives and SetupTimeOut are Debian-specific
//         compatibility aliases for this option.
//

func init() {
	funcInit := func() (res string) {
		// TODO
		panic("Not implemented")
		return
	}

	funcValid := func(value string) (res bool) {
		value = strings.TrimSpace(value)
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 0 {
			return false
		}
		return true
	}

	funcParse := func(input string) (values []string, err error) {
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

	sshInit(ServerAliveInterval, funcInit)
	sshValid(ServerAliveInterval, funcValid)
	sshParse(ServerAliveInterval, funcParse)
}
