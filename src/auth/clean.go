package auth

import "os"

// Clean removes what a session left: its lock, its temporary file, its
// log, its backup and the backup before it. A file already gone is no
// error: the session may have ended well and removed its own.
//
// The order matters little: each file stands alone, and none is read
// again once the session is over.
//
// Nothing here waits for the session to end: the caller does.
//
// Clean is safe to run twice.
func Clean(dir string) {
	os.Remove(dir + "/session.lock")
	os.Remove(dir + "/session.tmp")
	os.Remove(dir + "/session.log")
	for _, f := range []string{"/session.bak", "/session.bak.1"} {
		if err := os.Remove(dir + f); err != nil && !os.IsNotExist(err) {
			panic(err)
		}
	}
}
