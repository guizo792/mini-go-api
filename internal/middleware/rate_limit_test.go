package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimit_AllowsRequestsWithinLimit(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrapped := RateLimit(2, time.Minute)(next)

	for range 2 {
		req := httptest.NewRequest(http.MethodGet, "/user/orders", nil)
		req.RemoteAddr = "192.168.1.10:1234"
		rr := httptest.NewRecorder()

		wrapped.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
	}
}

func TestRateLimit_BlocksRequestsOverLimit(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrapped := RateLimit(1, time.Minute)(next)

	firstReq := httptest.NewRequest(http.MethodGet, "/user/orders", nil)
	firstReq.RemoteAddr = "10.0.0.1:1234"
	firstRR := httptest.NewRecorder()
	wrapped.ServeHTTP(firstRR, firstReq)

	if firstRR.Code != http.StatusOK {
		t.Fatalf("expected first request 200, got %d", firstRR.Code)
	}

	secondReq := httptest.NewRequest(http.MethodGet, "/user/orders", nil)
	secondReq.RemoteAddr = "10.0.0.1:1234"
	secondRR := httptest.NewRecorder()
	wrapped.ServeHTTP(secondRR, secondReq)

	if secondRR.Code != http.StatusTooManyRequests {
		t.Fatalf("expected second request 429, got %d", secondRR.Code)
	}

	if secondRR.Header().Get("Retry-After") == "" {
		t.Fatal("expected Retry-After header to be set")
	}
}
