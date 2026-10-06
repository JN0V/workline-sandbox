package auth

import "time"

// IdleTimeout signs a session out when nothing used it for this long.
const IdleTimeout = 30 * time.Minute

// Idle says whether a session last used at last is signed out at now.
func Idle(last, now time.Time) bool {
	idleFor := now.Sub(last)
	return idleFor >= IdleTimeout
}
