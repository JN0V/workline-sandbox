package auth

import (
	"testing"
	"time"
)

func TestLocked(t *testing.T) {
	now := time.Now()
	if Locked(4, now, now) {
		t.Error("four failures locked the account")
	}
	if !Locked(5, now, now) {
		t.Error("five failures did not lock the account")
	}
}
