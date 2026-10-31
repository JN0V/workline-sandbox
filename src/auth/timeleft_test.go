package auth

import (
	"testing"
	"time"
)

// Time left:
// Given a session last used 10 minutes ago,
// When I ask the time left,
// Then I get 20 minutes.
func TestTimeLeft(t *testing.T) {
	now := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	last := now.Add(-10 * time.Minute)

	got := TimeLeft(last, now)
	if got != 20*time.Minute {
		t.Fatalf("TimeLeft = %v, want %v", got, 20*time.Minute)
	}
}

// Already out:
// Given a session last used 40 minutes ago,
// When I ask the time left,
// Then I get 0.
func TestAlreadyOut(t *testing.T) {
	now := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	last := now.Add(-40 * time.Minute)

	got := TimeLeft(last, now)
	if got != 0 {
		t.Fatalf("TimeLeft = %v, want 0", got)
	}
}
