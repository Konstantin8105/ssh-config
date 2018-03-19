package ssh_config

import "testing"

// CanonicalizeFallbackLocal
//         Specifies whether to fail with an error when hostname canoni‐
//         calization fails.  The default, “yes”, will attempt to look up
//         the unqualified hostname using the system resolver's search
//         rules.  A value of “no” will cause ssh(1) to fail instantly if
//         CanonicalizeHostname is enabled and the target hostname cannot
//         be found in any of the domains specified by CanonicalDomains.
//

func TestCanonicalizeFallbackLocalInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[CanonicalizeFallbackLocal]
	if !ok {
		t.Errorf("Cannot found init-function")
	}
	if f == nil {
		t.Errorf("Init-function is nil")
	}
	if f() == "" {
		// TODO
		t.Errorf("Not acceptable empty init value")
	}
}

func TestCanonicalizeFallbackLocalValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[CanonicalizeFallbackLocal]
	if !ok {
		t.Errorf("Cannot found valid-function")
	}
	if f == nil {
		t.Errorf("Valid-function is nil")
	}
	if !f("") {
		// TODO
		t.Errorf("Not acceptable empty valid-value")
	}
}

func TestCanonicalizeFallbackLocalParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[CanonicalizeFallbackLocal]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
