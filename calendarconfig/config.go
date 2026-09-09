// Package calendarconfig provides strict config-compatible civil values.
//
// Deprecated: use github.com/faustbrian/go-calendar/adapters/config. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package calendarconfig

import (
	calendar "github.com/faustbrian/go-calendar"
	adapter "github.com/faustbrian/go-calendar/adapters/config"
)

// Date is a config ValueUnmarshaler for a required canonical civil date.
type Date struct{ value adapter.Date }

// NewDate wraps a calendar Date.
func NewDate(date calendar.Date) Date { return Date{value: adapter.NewDate(date)} }

// CalendarDate returns the decoded civil date.
func (d Date) CalendarDate() calendar.Date { return d.value.CalendarDate() }

// UnmarshalConfigValue accepts only a canonical date string.
func (d *Date) UnmarshalConfigValue(value any) error {
	if d == nil {
		return calendar.ErrInvalidDate
	}
	return d.value.UnmarshalConfigValue(value)
}

// MarshalText returns the canonical configuration value.
func (d Date) MarshalText() ([]byte, error) { return d.value.MarshalText() }

// UnmarshalText decodes a canonical configuration value.
func (d *Date) UnmarshalText(text []byte) error {
	if d == nil {
		return calendar.ErrInvalidDate
	}
	return d.value.UnmarshalText(text)
}
