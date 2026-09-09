// Package calendarvalidation provides dependency-neutral rules that can be
// wrapped by validation ValidatorFunc without coupling calendar core to it.
//
// Deprecated: use github.com/faustbrian/go-calendar/adapters/validation. This
// package remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package calendarvalidation

import (
	calendar "github.com/faustbrian/go-calendar"
	adapter "github.com/faustbrian/go-calendar/adapters/validation"
)

var (
	// ErrInvalidDate identifies an invalid Date value.
	ErrInvalidDate = adapter.ErrInvalidDate
	// ErrDateOutOfRange identifies a Date outside inclusive configured bounds.
	ErrDateOutOfRange = adapter.ErrDateOutOfRange
)

// Rule is a deterministic, side-effect-free date validation function.
type Rule func(calendar.Date) error

// ValidDate returns a rule that rejects the Date zero value.
func ValidDate() Rule {
	return Rule(adapter.ValidDate())
}

// DateRange returns an inclusive bounded date rule.
func DateRange(minimum, maximum calendar.Date) (Rule, error) {
	rule, err := adapter.DateRange(minimum, maximum)
	return Rule(rule), err
}
