package logic

import (
	"time"
)

type Sec float64

type timer struct {
	last time.Time
}

func NewTimer() *timer {
	return &timer{}
}

func (t *timer) Start() {
	t.last = time.Now()
}

func (t *timer) IsStarted() bool {
	return !t.last.IsZero()
}

func Touch(t *timer) time.Duration {
	if !t.IsStarted() {
		panic("timer hasnt been started")
	}

	now := time.Now()
	elpased := now.Sub(t.last)
	t.last = now
	return elpased
}
