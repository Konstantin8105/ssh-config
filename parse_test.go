package ssh_config

import (
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	files, err := filepath.Glob("testdata/*")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			// TODO
		})
	}
}
