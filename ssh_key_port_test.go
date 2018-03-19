package ssh_config

import (
	"fmt"
	"testing"
)

// Port    Specifies the port number to connect on the remote host.  The
//         default is 22.
//

func TestPortInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[Port]
	if !ok {
		t.Errorf("Cannot found init-function")
	}
	if f == nil {
		t.Errorf("Init-function is nil")
	}
	if f() == "" {
		t.Errorf("Not acceptable empty init value")
	}
	if f() != "22" {
		t.Errorf("Not acceptable post number : %v", f())
	}
}

func TestPortValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[Port]
	if !ok {
		t.Errorf("Cannot found valid-function")
	}
	if f == nil {
		t.Errorf("Valid-function is nil")
	}
	if !f(mapInit[Port]()) {
		t.Errorf("Not acceptable valid-value for init-value")
	}
	for i := 0; i < 65535+1; i++ {
		port := fmt.Sprintf(" %d    ", i)
		if !f(port) {
			t.Errorf("Error for valid-value : %s", port)
		}
	}
	if f(" -1   ") {
		t.Errorf("Error for NOT valid-value : -1")
	}
	if f("65536 ") {
		t.Errorf("Error for NOT valid-value : 65536")
	}
}

func TestPortParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[Port]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
