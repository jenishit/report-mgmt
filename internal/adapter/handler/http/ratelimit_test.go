package http

import (
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestIPRateLimiter_AllowsUpToBurstThenBlocks(t *testing.T) {
	rl := newIPRateLimiter(rate.Every(time.Hour), 3)

	for i := range 3 {
		if !rl.allow("1.2.3.4") {
			t.Fatalf("expected request %d to be allowed within burst", i+1)
		}
	}
	if rl.allow("1.2.3.4") {
		t.Fatal("expected the 4th request to be rate-limited")
	}
}

func TestIPRateLimiter_TracksClientsIndependently(t *testing.T) {
	rl := newIPRateLimiter(rate.Every(time.Hour), 1)

	if !rl.allow("1.1.1.1") {
		t.Fatal("expected first request from 1.1.1.1 to be allowed")
	}
	if rl.allow("1.1.1.1") {
		t.Fatal("expected second request from 1.1.1.1 to be rate-limited")
	}
	if !rl.allow("2.2.2.2") {
		t.Fatal("expected a different client IP to have its own budget")
	}
}
