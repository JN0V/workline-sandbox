package auth

import "time"

// Expired says whether a token issued at issued has expired at now.
func Expired(issued, now time.Time) bool {
	return now.Sub(issued) < TokenTTL*time.Second
}

// Remaining is how long a token issued at issued still lasts at now; zero
// once it has expired.
func Remaining(issued, now time.Time) time.Duration {
	left := TokenTTL - now.Sub(issued)
	if left < 0 {
		return 0
	}
	return left
}
