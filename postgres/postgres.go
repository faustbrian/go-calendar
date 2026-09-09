// Package postgres provides database/sql and pgx adapters for PostgreSQL date.
// Ordinary Date rejects NULL and infinity; InfinityDate models those sentinels
// explicitly when an application needs them.
//
// Deprecated: use github.com/faustbrian/go-calendar/adapters/postgres. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package postgres

import (
	"database/sql/driver"

	calendar "github.com/faustbrian/go-calendar"
	calendarpostgres "github.com/faustbrian/go-calendar/adapters/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	// ErrNull identifies SQL NULL where an ordinary civil date is required.
	ErrNull = calendarpostgres.ErrNull
	// ErrInfinity identifies PostgreSQL infinity where an ordinary date is required.
	ErrInfinity = calendarpostgres.ErrInfinity
)

// Date adapts a non-null, finite calendar.Date to database/sql and pgx.
type Date struct{ value calendarpostgres.Date }

// NewDate constructs a PostgreSQL adapter. Invalid dates remain invalid and
// return calendar.ErrInvalidDate when encoded.
func NewDate(date calendar.Date) Date { return Date{value: calendarpostgres.NewDate(date)} }

// CalendarDate returns the wrapped civil date.
func (d Date) CalendarDate() calendar.Date { return d.value.CalendarDate() }

// Value implements database/sql/driver.Valuer using canonical date text.
func (d Date) Value() (driver.Value, error) {
	return d.value.Value()
}

// Scan implements database/sql.Scanner.
func (d *Date) Scan(source any) error {
	if d == nil {
		return calendar.ErrInvalidDate
	}
	return d.value.Scan(source)
}

// DateValue implements pgtype.DateValuer.
func (d Date) DateValue() (pgtype.Date, error) {
	return d.value.DateValue()
}

// ScanDate implements pgtype.DateScanner.
func (d *Date) ScanDate(value pgtype.Date) error {
	if d == nil {
		return calendar.ErrInvalidDate
	}
	return d.value.ScanDate(value)
}

// InfinityKind classifies a finite or infinite PostgreSQL date.
type InfinityKind int8

const (
	// NegativeInfinity represents PostgreSQL -infinity.
	NegativeInfinity InfinityKind = -1
	// Finite represents an ordinary civil date.
	Finite InfinityKind = 0
	// PositiveInfinity represents PostgreSQL infinity.
	PositiveInfinity InfinityKind = 1
)

// InfinityDate is the explicit sum type for finite and infinite PostgreSQL
// dates. Its zero value is invalid because it has no finite Date.
type InfinityDate struct {
	value calendarpostgres.InfinityDate
}

// NewInfinityDate constructs an infinite value. Finite is rejected at encode
// time because callers must use NewFiniteDate with a concrete date.
func NewInfinityDate(kind InfinityKind) InfinityDate {
	return InfinityDate{value: calendarpostgres.NewInfinityDate(calendarpostgres.InfinityKind(kind))}
}

// NewFiniteDate constructs an infinity-aware finite value.
func NewFiniteDate(date calendar.Date) InfinityDate {
	return InfinityDate{value: calendarpostgres.NewFiniteDate(date)}
}

// Kind returns the value classification.
func (d InfinityDate) Kind() InfinityKind { return InfinityKind(d.value.Kind()) }

// Date returns the finite date, or an invalid Date for infinities.
func (d InfinityDate) Date() calendar.Date { return d.value.Date() }

// Value implements database/sql/driver.Valuer.
func (d InfinityDate) Value() (driver.Value, error) {
	return d.value.Value()
}

// Scan implements database/sql.Scanner.
func (d *InfinityDate) Scan(source any) error {
	if d == nil {
		return calendar.ErrInvalidDate
	}
	return d.value.Scan(source)
}

// DateValue implements pgtype.DateValuer.
func (d InfinityDate) DateValue() (pgtype.Date, error) {
	return d.value.DateValue()
}

// ScanDate implements pgtype.DateScanner.
func (d *InfinityDate) ScanDate(value pgtype.Date) error {
	if d == nil {
		return calendar.ErrInvalidDate
	}
	return d.value.ScanDate(value)
}
