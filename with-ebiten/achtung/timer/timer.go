package timer

import "time"

type Timer struct {
	lastUpdate time.Time
}

func NewTimer() Timer {
	return Timer{
		lastUpdate: time.Time{},
	}
}

func (t *Timer) Touch() time.Duration {
	now := time.Now()

	if t.lastUpdate.IsZero() {
		t.lastUpdate = now
	}

	elapsed := now.Sub(t.lastUpdate)
	t.lastUpdate = now
	return elapsed
}
