// Package calendarclock adapts a narrow wall-clock source to civil dates.
//
// Deprecated: use github.com/faustbrian/go-calendar/adapters/clock. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package calendarclock

import (
	"time"

	calendar "github.com/faustbrian/go-calendar"
	adapter "github.com/faustbrian/go-calendar/adapters/clock"
)

// ErrClockRequired identifies a missing wall-clock capability.
var ErrClockRequired = adapter.ErrClockRequired

// Clock is the exact wall-clock capability provided by clock.Clock.
type Clock interface {
	Now() time.Time
}

// Today obtains the civil date observed in an explicit location.
func Today(clock Clock, location *time.Location) (calendar.Date, error) {
	return adapter.Today(clock, location)
}
