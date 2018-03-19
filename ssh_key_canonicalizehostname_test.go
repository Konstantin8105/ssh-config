package ssh_config

import "testing"

// CanonicalizeHostname
//         Controls whether explicit hostname canonicalization is per‐
//         formed.  The default, “no”, is not to perform any name rewrit‐
//         ing and let the system resolver handle all hostname lookups.
//         If set to “yes” then, for connections that do not use a
//         ProxyCommand, ssh(1) will attempt to canonicalize the hostname
//         specified on the command line using the CanonicalDomains suf‐
//         fixes and CanonicalizePermittedCNAMEs rules.  If
//         CanonicalizeHostname is set to “always”, then canonicalization
//         is applied to proxied connections too.
//
//         If this option is enabled, then the configuration files are
//         processed again using the new target name to pick up any new
//         configuration in matching Host and Match stanzas.
//

func TestCanonicalizeHostnameInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[CanonicalizeHostname]
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

func TestCanonicalizeHostnameValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[CanonicalizeHostname]
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

func TestCanonicalizeHostnameParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[CanonicalizeHostname]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
