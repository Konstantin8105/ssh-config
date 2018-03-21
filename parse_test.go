package ssh_config

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Cannot continue parsing test : panic is not acceptable.")
		}
	}()

	files, err := filepath.Glob("testdata/*")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("Cannot parse '%v' : panic is not acceptable.",
						file)
				}
			}()

			contents, err := ioutil.ReadFile(file)
			if err != nil {
				t.Errorf("Cannot read file `%v`. err = %v", file, err)
				return
			}
			out, err := Parse(contents)
			if err != nil {
				t.Errorf("Cannot parse file `%v`. err = %v", file, err)
				return
			}
			if out == nil {
				t.Errorf("Output cannot by nil")
				return
			}
			// TODO : compare result
		})
	}
}
