package http

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// visitorWindow expired entries are pruned after this much inactivity.
const visitorWindow = 3 * time.Minute

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// ipRateLimiter is a per-client-IP token bucket limiter, used to slow down
// brute-force login/OTP-guessing attempts. It's in-memory, which is fine for
// a single instance; a multi-instance deployment would need a shared store
// (e.g. Redis) instead.
type ipRateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	r        rate.Limit
	burst    int
}

func newIPRateLimiter(r rate.Limit, burst int) *ipRateLimiter {
	rl := &ipRateLimiter{
		visitors: make(map[string]*visitor),
		r:        r,
		burst:    burst,
	}
	go rl.cleanupLoop()
	return rl
}

func (rl *ipRateLimiter) allow(key string) bool {
	rl.mu.Lock()
	v, exists := rl.visitors[key]
	if !exists {
		v = &visitor{limiter: rate.NewLimiter(rl.r, rl.burst)}
		rl.visitors[key] = v
	}
	v.lastSeen = time.Now()
	limiter := v.limiter
	rl.mu.Unlock()

	return limiter.Allow()
}

func (rl *ipRateLimiter) cleanupLoop() {
	for range time.Tick(time.Minute) {
		rl.mu.Lock()
		for key, v := range rl.visitors {
			if time.Since(v.lastSeen) > visitorWindow {
				delete(rl.visitors, key)
			}
		}
		rl.mu.Unlock()
	}
}

// rateLimitMiddleware limits requests per client IP to r requests/second
// with the given burst allowance.
func rateLimitMiddleware(r rate.Limit, burst int) gin.HandlerFunc {
	limiter := newIPRateLimiter(r, burst)

	return func(ctx *gin.Context) {
		if !limiter.allow(ctx.ClientIP()) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, newErrorResponse([]string{"too many requests, please try again later"}))
			return
		}
		ctx.Next()
	}
}
