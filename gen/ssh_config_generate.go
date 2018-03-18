package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type status int

const (
	newOption status = iota
	bodyOption
)

type option struct {
	name     string
	comments []string
}

func main() {
	cmd := exec.Command("man", "ssh_config")
	var sout bytes.Buffer
	var serr bytes.Buffer
	cmd.Stdout = &sout
	cmd.Stderr = &serr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Cannot run `man ssh_config`. err = %v\n stderr = %v",
			err, serr)
		return
	}
	lines := strings.Split(sout.String(), "\n")

	// Generate package comment
	filename := "../ssh_config.go"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Cannot create a file `%v`. err = %v",
			filename, err)
		return
	}
	defer func() {
		err = f.Sync()
		if err != nil {
			fmt.Println("Cannot sync file `%v`. err = %v",
				filename, err)
		}
		err = f.Close()
		if err != nil {
			fmt.Println("Cannot close file `%v`. err = %v",
				filename, err)
		}
	}()

	for _, line := range lines {
		_, err = f.Write([]byte("// "))
		if err != nil {
			fmt.Println("Cannot write to file `%v`. err = %v",
				filename, err)
			return
		}
		_, err = f.WriteString(line)
		if err != nil {
			fmt.Println("Cannot write to file `%v`. err = %v",
				filename, err)
			return
		}
		_, err = f.Write([]byte("\n"))
		if err != nil {
			fmt.Println("Cannot write to file `%v`. err = %v",
				filename, err)
			return
		}
	}
	_, err = f.WriteString("package ssh_config\n")
	if err != nil {
		fmt.Println("Cannot write to file `%v`. err = %v",
			filename, err)
		return
	}

	// Parse `man documentation` to golang struct
	paragrahName := "DESCRIPTION"
	foundParagrah := false
	indentParagrah := 5
	firstOption := "Host"
	foundOptions := false

	var st status = newOption
	options := make([]option, 0, 1)
	var presentOption option

	for _, line := range lines {
		if len(line) > 0 && line[0] != ' ' {
			// found paragraf
			if foundParagrah {
				foundParagrah = false
			}
			if line == paragrahName {
				foundParagrah = true
			}
		}
		if !foundParagrah {
			continue
		}
		if !((len(line) >= indentParagrah && len(strings.TrimSpace(line[:indentParagrah])) == 0) ||
			(len(line) == 0)) {
			continue
		}
		if !foundOptions {
			s := strings.Split(strings.TrimSpace(line), " ")
			if len(s) == 0 {
				continue
			}
			if s[0] != firstOption {
				continue
			} else {
				foundOptions = true
			}
		}
		if len(strings.TrimSpace(line)) == 0 {
			// for empty line
			line = ""
		}
		if len(line) > 0 {
			line = line[indentParagrah:]
		}

		if len(line) > 0 && line[0] != ' ' {
			st = newOption
		} else {
			st = bodyOption
		}

		if st == newOption {
			options = append(options, presentOption)
			presentOption = option{}
			// search name of option
			s := strings.Split(strings.TrimSpace(line), " ")
			if len(s) == 0 {
				panic(fmt.Errorf("Wrong string line with name : %s", line))
			}
			presentOption.name = s[0]
		}
		presentOption.comments = append(presentOption.comments, line)
	}
	options = append(options, presentOption)
	options = options[1:]

	for _, opt := range options {
		fmt.Printf("%#v\n", opt.name)
	}
}
