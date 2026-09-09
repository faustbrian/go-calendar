// Package calendartemporal converts civil boundaries into values suitable for
// temporal without adding interval algebra to calendar core.
//
// Deprecated: use github.com/faustbrian/go-calendar/adapters/temporal. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package calendartemporal

import (
	"time"

	calendar "github.com/faustbrian/go-calendar"
	adapter "github.com/faustbrian/go-calendar/adapters/temporal"
	calendartz "github.com/faustbrian/go-calendar/timezone"
)

var (
	// ErrReversed identifies civil endpoints in descending order.
	ErrReversed = adapter.ErrReversed
	// ErrRangeLimit identifies a date sequence exceeding its explicit bound.
	ErrRangeLimit = adapter.ErrRangeLimit
)

// InstantRange is a canonical start-inclusive, end-exclusive adapter value.
type InstantRange struct{ value adapter.InstantRange }

// Start returns the inclusive instant boundary.
func (r InstantRange) Start() time.Time { return r.value.Start() }

// End returns the exclusive instant boundary.
func (r InstantRange) End() time.Time { return r.value.End() }

// Includes reports membership using [start,end) semantics.
func (r InstantRange) Includes(value time.Time) bool { return r.value.Includes(value) }

// InclusiveDates converts inclusive civil endpoints to an exclusive instant
// range. Its boundaries can be passed directly to temporal/instant.Range.
func InclusiveDates(first, last calendar.Date, location *time.Location, policy calendartz.Resolution) (InstantRange, error) {
	value, err := adapter.InclusiveDates(first, last, location, policy)
	return InstantRange{value: value}, err
}

// Sequence returns every date in the inclusive range with explicit work bound.
func Sequence(first, last calendar.Date, limit int) ([]calendar.Date, error) {
	return adapter.Sequence(first, last, limit)
}
