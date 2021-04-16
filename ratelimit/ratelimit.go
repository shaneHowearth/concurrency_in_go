// Package ratelimit -
package ratelimit

import (
	"context"
	"sort"
	"time"
)

// Open -
func Open() *APIConnection {
	secondLimit := NewLimiter(Per(2, time.Second), 1)
	minuteLimit := NewLimiter(Per(10, time.Minute), 10)
	return &APIConnection{
		rateLimiter: MultiLimiter(secondLimit, minuteLimit),
	}

}

// APIConnection -
type APIConnection struct {
	rateLimiter RateLimiter
}

// ReadFile -
func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err

	}
	// Pretend we do work here
	return nil

}

// ResolveAddress -
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err

	}
	// Pretend we do work here
	return nil

}

// Limit defines the maximum frequency of some events. Limit is
// represented as number of events per second. A zero Limit allows no
// events.
type Limit float64

// Limiter -
type Limiter struct{}

// NewLimiter returns a new Limiter that allows events up to rate r
// and permits bursts of at most b tokens.
func NewLimiter(r Limit, b int) *Limiter {
	return nil
}

// Limit -
func (lim *Limiter) Limit() Limit {
	return Limit(0)
}

// Every converts a minimum time interval between events to a Limit.
func Every(interval time.Duration) Limit {
	return Limit(0)
}

// Per -
func Per(eventCount int, duration time.Duration) Limit {
	return Every(duration / time.Duration(eventCount))

}

// Wait is shorthand for WaitN(ctx, 1).
func (lim *Limiter) Wait(ctx context.Context) error { return nil }

// WaitN blocks until lim permits n events to happen.
// It returns an error if n exceeds the Limiter's burst size, the Context is
// canceled, or the expected wait time exceeds the Context's Deadline.
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error) {
	return
}

// RateLimiter -
type RateLimiter interface {
	Wait(context.Context) error
	Limit() Limit
}

// MultiLimiter -
// ignore return unexported type linting error
// nolint:golint
func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()

	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err

		}

	}
	return nil

}
func (l *multiLimiter) Limit() Limit {
	return l.limiters[0].Limit()

}
