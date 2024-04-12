package exponential_backoff

import (
	"math"
	"time"

	lsmath "github.com/cuongpiger/joat/math"
)

type BackOff struct {
	Revert                        bool
	ExponentialBackoffCeilingSecs int64
	Timeout                       time.Duration

	attempts int
}

func NewBackOff(psteps int, pebcs int64, prev bool, pto time.Duration) *BackOff {
	attempts := 0
	if prev {
		attempts = psteps
	}

	return &BackOff{
		Revert:                        prev,
		ExponentialBackoffCeilingSecs: pebcs,
		Timeout:                       pto,
		attempts:                      attempts,
	}
}

func (s *BackOff) Step() time.Duration {
	s.attemp()
	return s.delay()
}

func (s *BackOff) delay() time.Duration {
	if s.attempts < 1 {
		return time.Duration(s.ExponentialBackoffCeilingSecs) * time.Second
	}

	delaySecs := int64(math.Floor((math.Pow(2, float64(s.attempts)) - 1) * 0.5))
	if s.Revert {
		return time.Duration(lsmath.MaxNumeric(s.ExponentialBackoffCeilingSecs, delaySecs)) * time.Second
	} else {
		return time.Duration(lsmath.MinNumeric(s.ExponentialBackoffCeilingSecs, delaySecs)) * time.Second
	}
}

func (s *BackOff) attemp() {
	if s.Revert {
		s.attempts--
	} else {
		s.attempts++
	}

	if s.attempts < 0 {
		s.attempts = 1
	}
}
