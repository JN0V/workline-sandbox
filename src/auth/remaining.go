package auth

// Remaining is how many seconds a token issued at issued still lasts at now.
func Remaining(issued, now int64) int64 {
	left := TokenTTL - (now - issued)
	if left < 0 {
		return 0
	}
	return left
}

func Expired(issued, now int64) bool { return Remaining(issued, now) == 0 }
