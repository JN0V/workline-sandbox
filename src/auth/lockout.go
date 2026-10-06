package auth

import "time"

// MaxFailures is how many failed sign-ins in a row lock an account.
const MaxFailures = 5

// LockFor is how long a locked account stays locked.
const LockFor = 15 * time.Minute

// Locked says whether an account whose last failures in a row number
// failures, the last of them at last, is locked at now.
func Locked(failures int, last, now time.Time) bool {
	return failures >= MaxFailures && now.Sub(last) < LockFor
}
