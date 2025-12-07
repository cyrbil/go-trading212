package trading212

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"time"
)

const (
	// RateLimitHeaderLimit x-ratelimit-limit: The total number of requests allowed in the current time period.
	RateLimitHeaderLimit = "x-ratelimit-limit"
	// RateLimitHeaderPeriod x-ratelimit-period: The duration of the time period in seconds.
	RateLimitHeaderPeriod = "x-ratelimit-period"
	// RateLimitHeaderRemaining x-ratelimit-remaining: The number of requests you have left in the current period.
	RateLimitHeaderRemaining = "x-ratelimit-remaining"
	// RateLimitHeaderReset x-ratelimit-reset: A Unix timestamp indicating the exact time when the limit will be reset.
	RateLimitHeaderReset = "x-ratelimit-reset"
	// RateLimitHeaderUsed x-ratelimit-used: The number of requests you have already made in the current period.
	RateLimitHeaderUsed = "x-ratelimit-used"
)

var (
	errHeaderNotFound   = errors.New("header not found")
	errHeaderConversion = errors.New("error conversion header")
)

func headerNotFoundError(header string) error {
	return fmt.Errorf("%w: %v", errHeaderNotFound, header)
}

func headerConversionError(header string, value string) error {
	return fmt.Errorf("%w: header %v with value %v is not an integer", errHeaderConversion, header, value)
}

// APIRateLimits rate-limits.
type APIRateLimits struct {
	// Limit x-ratelimit-limit: The total number of requests allowed in the current time period.
	Limit uint64
	// Period x-ratelimit-period: The duration of the time period in seconds.
	Period time.Duration
	// Remaining x-ratelimit-remaining: The number of requests you have left in the current period.
	Remaining uint64
	// Reset x-ratelimit-reset: A Unix timestamp indicating the exact time when the limit will be fully reset.
	Reset time.Time
	// Used x-ratelimit-used: The number of requests you have already made in the current period.
	Used uint64
}

// RateLimiter type
type RateLimiter struct {
	limits map[string]APIRateLimits
	sleep  func(time.Duration)
}

// NewRateLimiter creates a RateLimiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limits: make(map[string]APIRateLimits),
		sleep:  time.Sleep,
	}
}

// ApplyRateLimit will sleep if a rate limit is in place.
func (r *RateLimiter) ApplyRateLimit(path string) {
	limits, found := r.limits[path]
	if !found {
		return
	}

	slog.Debug("Limit rate", "limits", limits)

	if limits.Remaining > 0 {
		return
	}

	now := time.Now()
	if now.After(limits.Reset) {
		return
	}

	r.sleep(time.Until(limits.Reset))
}

// ParseRateLimits parses the http response rate limit headers.
func (r *RateLimiter) ParseRateLimits(path string, response *http.Response) error {
	if response == nil || response.Header == nil || response.Request == nil || response.Request.URL == nil {
		return headerNotFoundError("response is nil")
	}
	headers := map[string]uint64{
		RateLimitHeaderLimit:     0,
		RateLimitHeaderPeriod:    0,
		RateLimitHeaderRemaining: 0,
		RateLimitHeaderReset:     0,
		RateLimitHeaderUsed:      0,
	}

	for header := range headers {
		str := response.Header.Get(header)
		if str == "" {
			return headerNotFoundError(header)
		}

		value, err := strconv.ParseUint(str, 10, 0)
		if err != nil {
			return headerConversionError(header, str)
		}

		headers[header] = value
	}

	if headers[RateLimitHeaderPeriod] > math.MaxInt64 {
		return errHeaderConversion
	}

	if headers[RateLimitHeaderReset] > math.MaxInt64 {
		return errHeaderConversion
	}

	rateLimits := &APIRateLimits{
		Limit:     headers[RateLimitHeaderLimit],
		Period:    time.Duration(headers[RateLimitHeaderPeriod]) * time.Second, //nolint:gosec
		Remaining: headers[RateLimitHeaderRemaining],
		Reset:     time.Unix(int64(headers[RateLimitHeaderReset]), 0), //nolint:gosec
		Used:      headers[RateLimitHeaderUsed],
	}

	r.limits[path] = *rateLimits
	return nil
}
