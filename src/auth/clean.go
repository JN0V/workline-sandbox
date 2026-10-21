package auth

import "os"

// Clean removes what a session left.
func Clean(dir string) {
	os.Remove(dir + "/session.lock")
	os.Remove(dir + "/session.tmp")
	os.Remove(dir + "/session.log")
}
