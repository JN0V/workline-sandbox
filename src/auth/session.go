package auth

import "time"

// Expired says whether a token issued at issued has expired at now.
func Expired(issued, now time.Time) bool {
	return now.Sub(issued) < TokenTTL*time.Second
}
