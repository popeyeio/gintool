package limiter

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/popeyeio/handy"
	"golang.org/x/time/rate"
)

var (
	ErrFormat = errors.New("limiter format error")
	ErrPeriod = errors.New("period error")

	periods = map[string]time.Duration{
		"S": time.Second,
		"M": time.Minute,
		"H": time.Hour,
	}
)

type TokenLimiter struct {
	limiter *rate.Limiter
}

var _ Limiter = (*TokenLimiter)(nil)

// if rate is 10 per second and burst is 20, the format is "10-S-20".
func NewTokenLimiter(format string) (*TokenLimiter, error) {
	tokens := strings.Split(format, handy.StrHyphen)
	if len(tokens) != 3 {
		return nil, ErrFormat
	}

	period, exists := periods[strings.ToUpper(tokens[1])]
	if !exists {
		return nil, ErrPeriod
	}

	burst, err := strconv.ParseInt(tokens[2], 10, 64)
	if err != nil {
		return nil, err
	}

	number, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		return nil, err
	}

	limit := rate.Inf
	if number >= 0 {
		limit = rate.Limit(float64(number) / period.Seconds())
	}

	return &TokenLimiter{
		limiter: rate.NewLimiter(limit, int(burst)),
	}, nil
}

func (l *TokenLimiter) Allow() bool {
	return l.limiter.Allow()
}

func (l *TokenLimiter) AllowN(n int) bool {
	return l.limiter.AllowN(time.Now(), n)
}
