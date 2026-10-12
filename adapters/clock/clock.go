// Package calendarclock adapts a caller-owned wall-clock source to civil dates.
// Today performs one synchronous call and owns no cancellation or clock lifecycle.
// Supply an explicit non-nil location and a bounded usable clock implementation.
package calendarclock

import (
	"errors"
	"time"

	calendar "github.com/faustbrian/go-calendar/v2"
)

// ErrClockRequired identifies a missing wall-clock capability.
var ErrClockRequired = errors.New("calendar/clock: clock required")

// Clock is the exact wall-clock capability provided by clock.Clock.
type Clock interface {
	// Now returns a bounded synchronous instant; caller owns concurrency and panic behavior.
	Now() time.Time
}

// Today calls Now once, then extracts civil components in an explicit location.
// Nil clock interfaces fail; typed-nil/custom callback behavior is not recovered.
// A nil location is rejected after the clock call. No deadline is imposed.
func Today(clock Clock, location *time.Location) (calendar.Date, error) {
	if clock == nil {
		return calendar.Date{}, ErrClockRequired
	}
	return calendar.DateFromTime(clock.Now(), location)
}
