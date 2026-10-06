package auth

import (
	"testing"
	"time"
)

func TestIdle(t *testing.T) {
	last := time.Date(2026, 10, 6, 12, 0, 0, 0, time.UTC)
	if Idle(last, last.Add(29*time.Minute)) {
		t.Error("29 minutes: signed out")
	}
	if !Idle(last, last.Add(31*time.Minute)) {
		t.Error("31 minutes: still signed in")
	}
}
