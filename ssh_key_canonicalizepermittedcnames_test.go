package ssh_config

import "testing"

// CanonicalizePermittedCNAMEs
//         Specifies rules to determine whether CNAMEs should be followed
//         when canonicalizing hostnames.  The rules consist of one or
//         more arguments of source_domain_list:target_domain_list, where
//         source_domain_list is a pattern-list of domains that may follow
//         CNAMEs in canonicalization, and target_domain_list is a pat‐
//         tern-list of domains that they may resolve to.
//
//         For example, “*.a.example.com:*.b.example.com,*.c.example.com”
//         will allow hostnames matching “*.a.example.com” to be canoni‐
//         calized to names in the “*.b.example.com” or “*.c.example.com”
//         domains.
//

func TestCanonicalizePermittedCNAMEsInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[CanonicalizePermittedCNAMEs]
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

func TestCanonicalizePermittedCNAMEsValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[CanonicalizePermittedCNAMEs]
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

func TestCanonicalizePermittedCNAMEsParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[CanonicalizePermittedCNAMEs]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
