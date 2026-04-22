package middleware

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type clientWindow struct {
	count       int
	windowStart time.Time
}

func clientIP(r *http.Request) string {
	xff := strings.TrimSpace(r.Header.Get("X-Forwarded-For"))
	if xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func RateLimit(limit int, window time.Duration) func(http.Handler) http.Handler {
	if limit <= 0 {
		limit = 1
	}
	if window <= 0 {
		window = time.Minute
	}

	var mu sync.Mutex
	clients := make(map[string]*clientWindow)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			ip := clientIP(r)

			mu.Lock()
			entry, ok := clients[ip]
			if !ok {
				entry = &clientWindow{count: 0, windowStart: now}
				clients[ip] = entry
			}

			if now.Sub(entry.windowStart) >= window {
				entry.count = 0
				entry.windowStart = now
			}

			if entry.count >= limit {
				retryAfter := int(window.Seconds() - now.Sub(entry.windowStart).Seconds())
				if retryAfter < 1 {
					retryAfter = 1
				}
				mu.Unlock()

				w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			entry.count++
			mu.Unlock()

			next.ServeHTTP(w, r)
		})
	}
}
