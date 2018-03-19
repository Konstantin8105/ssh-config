package ssh_config

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestAmountTodo(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	var amount int
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}

		amount += bytes.Count(content, []byte("TODO"))
	}

	if amount > 0 {
		t.Errorf("Amount of words `TODO` is not zero. Amount = %v", amount)
	}
}
