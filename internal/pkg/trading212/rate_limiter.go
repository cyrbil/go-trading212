package trading212

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type APIRateLimits struct {
	// Limit x-ratelimit-limit: The total number of requests allowed in the current time period.
	Limit uint
	// Period x-ratelimit-period: The duration of the time period in seconds.
	Period time.Duration
	// Remaining x-ratelimit-remaining: The number of requests you have left in the current period.
	Remaining uint
	// Reset x-ratelimit-reset: A Unix timestamp indicating the exact time when the limit will be fully reset.
	Reset time.Time
	// Used x-ratelimit-used: The number of requests you have already made in the current period.
	Used uint
}

// ParseRateLimits parses the http response rate limit headers
func ParseRateLimits(response *http.Response) (*APIRateLimits, error) {
	prefix := "x-ratelimit"
	headers := map[string]uint{
		"limit":     0,
		"period":    0,
		"remaining": 0,
		"reset":     0,
		"used":      0,
	}

	for key := range headers {
		header := fmt.Sprintf("%s-%s", prefix, key)
		str := response.Header.Get(header)
		if str == "" {
			return nil, fmt.Errorf("%s header not found", header)
		}

		value, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("error converting header %v value %v to integer", header, str)
		}
		headers[key] = uint(value)
	}

	rateLimits := &APIRateLimits{
		Limit:     headers["limit"],
		Period:    time.Duration(headers["period"]) * time.Second,
		Remaining: headers["remaining"],
		Reset:     time.Unix(int64(headers["reset"]), 0),
		Used:      headers["used"],
	}

	return rateLimits, nil
}

// ApplyRateLimit will sleep if a rate limit is in place
func ApplyRateLimit(path string, rateLimits map[string]APIRateLimits) {
	limits, found := rateLimits[path]
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
	time.Sleep(time.Until(limits.Reset))
}
