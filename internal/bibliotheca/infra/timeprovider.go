package infra

import "time"

var _ TimeProvider = &timeProvider{}

type TimeProvider interface {
	Now() time.Time
}

type timeProvider struct{}

func NewTimeProvider() TimeProvider {
	return &timeProvider{}
}

func (t timeProvider) Now() time.Time {
	return time.Now().UTC()
}
