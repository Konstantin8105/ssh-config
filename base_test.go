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
		if file == "base_test.go" {
			continue
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		a := bytes.Count(content, []byte("TODO"))
		if a > 0 {
			t.Logf("Found %d `TODO` tags in file : %s", a, file)
		}
		amount += a
	}
	if amount > 0 {
		t.Errorf("Amount of words `TODO` is not zero. Amount = %v", amount)
	}
}

func TestAmountFix(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	var amount int
	for _, file := range files {
		if file == "base_test.go" {
			continue
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		a := bytes.Count(content, []byte("FIX"))
		if a > 0 {
			t.Logf("Found %d `FIX` tags in file : %s", a, file)
		}
		amount += a
	}
	if amount > 0 {
		t.Errorf("Amount of words `FIX` is not zero. Amount = %v", amount)
	}
}

func TestAmountPanic(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	var amount int
	for _, file := range files {
		if file == "base_test.go" {
			continue
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		a := bytes.Count(content, []byte("panic("))
		if a > 0 {
			t.Logf("Found %d `panic` tags in file : %s", a, file)
		}
		amount += a
	}
	if amount > 0 {
		t.Errorf("Amount of words `panic` is not zero. Amount = %v", amount)
	}
}
