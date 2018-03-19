package ssh_config

import "testing"

// ServerAliveCountMax
//         Sets the number of server alive messages (see below) which may
//         be sent without ssh(1) receiving any messages back from the
//         server.  If this threshold is reached while server alive mes‐
//         sages are being sent, ssh will disconnect from the server, ter‐
//         minating the session.  It is important to note that the use of
//         server alive messages is very different from TCPKeepAlive
//         (below).  The server alive messages are sent through the
//         encrypted channel and therefore will not be spoofable.  The TCP
//         keepalive option enabled by TCPKeepAlive is spoofable.  The
//         server alive mechanism is valuable when the client or server
//         depend on knowing when a connection has become inactive.
//
//         The default value is 3.  If, for example, ServerAliveInterval
//         (see below) is set to 15 and ServerAliveCountMax is left at the
//         default, if the server becomes unresponsive, ssh will discon‐
//         nect after approximately 45 seconds.
//

func TestServerAliveCountMaxInitFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapInit[ServerAliveCountMax]
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

func TestServerAliveCountMaxValidFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapValid[ServerAliveCountMax]
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

func TestServerAliveCountMaxParseFunc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic in test is not acceptable")
		}
	}()
	f, ok := mapParse[ServerAliveCountMax]
	if !ok {
		t.Errorf("Cannot found parse-function")
	}
	if f == nil {
		t.Errorf("Parse-function is nil")
	}
	// TODO
}
