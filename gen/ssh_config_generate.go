package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

type status int

const (
	newOption status = iota
	bodyOption
)

type Option struct {
	Name     string
	Comments []string
}

var doNotEdit string = "\n// Code generated automatically. DO NOT EDIT.\n\n"

func main() {
	options, err := getManOptions()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = generateConstants(options)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = generateConvert(options)
	if err != nil {
		fmt.Println(err)
		return
	}
	/* DO NOT UNCOMMENT
	for _, opt := range options {
		err = generatePrepareSourceForKeys(opt)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	*/
}

func getManOptions() (options []Option, err error) {
	env := os.Getenv("MANWIDTH")
	err = os.Setenv("MANWIDTH", "78")
	if err != nil {
		fmt.Printf("Cannot set variable environment. err = ", err)
		return
	}
	defer func() {
		_ = os.Setenv("MANWIDTH", env)
	}()
	cmd := exec.Command("man", "ssh_config")
	var sout bytes.Buffer
	var serr bytes.Buffer
	cmd.Stdout = &sout
	cmd.Stderr = &serr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Cannot run `man ssh_config`. err = %v\n stderr = %v",
			err, serr)
		return
	}
	lines := strings.Split(sout.String(), "\n")

	// Generate package comment
	filename := "doc.go"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Cannot create a file `%v`. err = %v",
			filename, err)
		return
	}
	defer func() {
		err = f.Sync()
		if err != nil {
			fmt.Printf("Cannot sync file `%v`. err = %v",
				filename, err)
		}
		err = f.Close()
		if err != nil {
			fmt.Printf("Cannot close file `%v`. err = %v",
				filename, err)
		}
	}()

	for _, line := range lines {
		_, err = f.Write([]byte("// "))
		if err != nil {
			fmt.Printf("Cannot write to file `%v`. err = %v",
				filename, err)
			return
		}
		_, err = f.WriteString(line)
		if err != nil {
			fmt.Printf("Cannot write to file `%v`. err = %v",
				filename, err)
			return
		}
		_, err = f.Write([]byte("\n"))
		if err != nil {
			fmt.Printf("Cannot write to file `%v`. err = %v",
				filename, err)
			return
		}
	}
	_, err = f.WriteString("package ssh_config\n")
	if err != nil {
		fmt.Printf("Cannot write to file `%v`. err = %v",
			filename, err)
		return
	}

	f.WriteString(doNotEdit)

	// Parse `man documentation` to golang struct
	paragrahName := "DESCRIPTION"
	foundParagrah := false
	indentParagrah := 5
	firstOption := "Host"
	foundOptions := false

	var st status = newOption
	var presentOption Option

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
			presentOption = Option{}
			// search name of option
			s := strings.Split(strings.TrimSpace(line), " ")
			if len(s) == 0 {
				panic(fmt.Errorf("Wrong string line with name : %s", line))
			}
			presentOption.Name = s[0]
		}
		presentOption.Comments = append(presentOption.Comments, line)
	}
	options = append(options, presentOption)
	options = options[1:]
	return
}

func generateConstants(options []Option) (err error) {
	// generate golang constants
	filename := "constants.go"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Cannot create a file `%v`. err = %v",
			filename, err)
		return
	}
	defer func() {
		err = f.Sync()
		if err != nil {
			fmt.Printf("Cannot sync file `%v`. err = %v",
				filename, err)
		}
		err = f.Close()
		if err != nil {
			fmt.Printf("Cannot close file `%v`. err = %v",
				filename, err)
		}
	}()
	src := `
{{ range .Comments }}// {{ . }}
{{ end }}const {{ .Name }} SSHKey = "{{ .Name }}"
	`
	t := template.Must(template.New("src").Parse(src))

	var bufStruct bytes.Buffer
	for _, opt := range options {
		err = t.Execute(&bufStruct, opt)
		if err != nil {
			fmt.Println("executing template:", err)
			return
		}
	}

	f.WriteString("package ssh_config\n")
	f.WriteString(doNotEdit)
	f.WriteString("type SSHKey string\n")

	_, err = f.WriteString(bufStruct.String())
	if err != nil {
		fmt.Printf("Cannot write to file `%v`. err = %v",
			filename, err)
		return
	}
	return
}

func generateConvert(options []Option) (err error) {
	filename := "convert.go"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Cannot create a file `%v`. err = %v",
			filename, err)
		return
	}
	defer func() {
		err = f.Sync()
		if err != nil {
			fmt.Printf("Cannot sync file `%v`. err = %v",
				filename, err)
		}
		err = f.Close()
		if err != nil {
			fmt.Printf("Cannot close file `%v`. err = %v",
				filename, err)
		}
	}()
	src := `
	case "{{ .Name }}":
		return {{ .Name }}, nil
`
	t := template.Must(template.New("src").Parse(src))

	var bufStruct bytes.Buffer
	for _, opt := range options {
		err = t.Execute(&bufStruct, opt)
		if err != nil {
			fmt.Println("executing template:", err)
			return
		}
	}

	f.WriteString("package ssh_config\n")
	f.WriteString(doNotEdit)
	f.WriteString(`

import "fmt"

// Convert return ssh-key if input name is valid, or error
func Convert(name string) (key SSHKey, err error) {
	switch name {
`)
	_, err = f.WriteString(bufStruct.String())
	if err != nil {
		fmt.Printf("Cannot write to file `%v`. err = %v",
			filename, err)
		return
	}
	f.WriteString(`
	}
	return key, fmt.Errorf("Not valid name of ssh key : %s", name)
}`)
	return
}

func generatePrepareSourceForKeys(sshkey Option) (err error) {
	filename := "ssh_key_" + strings.ToLower(sshkey.Name) + ".go"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Cannot create a file `%v`. err = %v",
			filename, err)
		return
	}
	defer func() {
		err = f.Sync()
		if err != nil {
			fmt.Printf("Cannot sync file `%v`. err = %v",
				filename, err)
		}
		err = f.Close()
		if err != nil {
			fmt.Printf("Cannot close file `%v`. err = %v",
				filename, err)
		}
	}()

	f.WriteString("package ssh_config\n")

	for _, c := range sshkey.Comments {
		f.WriteString("// " + c + "\n")
	}

	f.WriteString(fmt.Sprintf(`

func init(){
	funcInit := func()(res string){
		// TODO
		return
	}

	funcValid := func(value string) (res bool){
		// TODO
		return
	}

	funcParse := func(input string) (values []string, err error){
		// TODO
		return
	}

	sshInit(%s, funcInit)
	sshValid(%s, funcValid)
	sshParse(%s, funcParse)
}
`, sshkey.Name, sshkey.Name, sshkey.Name))

	return
}
