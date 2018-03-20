package ssh_config

import (
	"bytes"
	"fmt"
)

type itemType int

const (
	itemEmpty itemType = iota
	itemComment
	itemSshKey
)

type block struct {
	it     itemType
	sshkey []byte
	value  []byte
}

// Parse file to ssh structure
func Parse(content []byte) (out []block, err error) {
	lines := bytes.Split(content, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line = bytes.TrimSpace(line)
		// Empty line
		if len(line) == 0 {
			out = append(out, block{
				it: itemEmpty,
			})
			continue
		}
		// Comment
		if line[0] == '#' {
			out = append(out, block{
				it:    itemComment,
				value: line[1:],
			})
			continue
		}
		// SshKey
	again:
		if line[len(line)-1] == ',' {
			if i >= len(lines)-1 {
				err = fmt.Errorf("Not correct comma at the end of line : `%v`",
					line)
				return
			}
			i++
			line = append(line, lines[i]...)
			goto again
		}
		var found bool
		for j := 0; j < len(line); j++ {
			switch line[j] {
			case ' ', '\n', '\x00', '\r':
				out = append(out, block{
					it:     itemSshKey,
					sshkey: line[:j],
					value:  line[j+1:],
				})
				found = true
				j = len(line)
			}
		}
		if found {
			continue
		}
		out = append(out, block{
			it:     itemSshKey,
			sshkey: line,
		})
	}
	for _, o := range out {
		if o.it == itemComment || o.it == itemEmpty {
			continue
		}
		fmt.Printf("> |%v|%v|%v|\n",
			o.it,
			string(o.sshkey),
			string(o.value))
		var sshkey SSHKey
		sshkey, err = Convert(string(o.sshkey))
		if err != nil {
			return
		}
		var values []string
		values, err = mapParse[sshkey](string(o.value))
		if err != nil {
			return
		}
		for _, v := range values {
			if !mapValid[sshkey](v) {
				err = fmt.Errorf("not valid value %v : %v ",
					string(o.sshkey), string(v))
				return
			}
		}
	}
	return
}
