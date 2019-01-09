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
	ErrRateFormat       = errors.New("rate format error")
	ErrLimitNotPositive = errors.New("limit not positive")
	ErrPeriod           = errors.New("period error")

	periods = map[string]time.Duration{
		"S": time.Second,
		"M": time.Minute,
		"H": time.Hour,
	}
)

type Limiter interface {
	Allow() bool
	AllowN(int) bool
}

type TokenLimiter struct {
	limiter *rate.Limiter
}

var _ Limiter = (*TokenLimiter)(nil)

func NewTokenLimiter(format string, burst int) (*TokenLimiter, error) {
	tokens := strings.Split(format, handy.StrHyphen)
	if len(tokens) != 2 {
		return nil, ErrRateFormat
	}

	limit, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		return nil, err
	}
	if limit <= 0 {
		return nil, ErrLimitNotPositive
	}

	period, exists := periods[strings.ToUpper(tokens[1])]
	if !exists {
		return nil, ErrPeriod
	}

	return &TokenLimiter{
		limiter: rate.NewLimiter(rate.Limit(1/period.Seconds()), burst),
	}, nil
}

func (l *TokenLimiter) Allow() bool {
	return l.limiter.Allow()
}

func (l *TokenLimiter) AllowN(n int) bool {
	return l.limiter.AllowN(time.Now(), n)
}
