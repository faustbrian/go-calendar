// Package calendarwire provides bounded canonical wire helpers for Date.
//
// Deprecated: use github.com/faustbrian/go-calendar/adapters/wire. This package
// remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package calendarwire

import (
	"cmp"
	"slices"

	calendar "github.com/faustbrian/go-calendar"
	adapter "github.com/faustbrian/go-calendar/adapters/wire"
)

const (
	// Version identifies the stable canonical date wire contract.
	Version = adapter.Version
	// MaxBytes bounds an encoded canonical date including JSON syntax.
	MaxBytes = adapter.MaxBytes
)

// ErrSizeLimit identifies input exceeding MaxBytes.
var ErrSizeLimit = adapter.ErrSizeLimit

// EncodeDate encodes a Date as the version-1 canonical JSON string. The
// compatibility implementation remains local because another package-call
// boundary changes its published allocation budget.
func EncodeDate(date calendar.Date) ([]byte, error) {
	if !date.IsValid() {
		return nil, calendar.ErrInvalidDate
	}
	return date.MarshalJSON()
}

// DecodeDate decodes exactly one bounded canonical JSON date string.
func DecodeDate(payload []byte) (calendar.Date, error) {
	if cmp.Compare(len(payload), MaxBytes) == 1 {
		return calendar.Date{}, ErrSizeLimit
	}
	if len(payload) != calendar.MaxParseBytes+2 {
		return calendar.Date{}, calendar.ErrInvalidFormat
	}
	if slices.Contains([]bool{payload[0] == '"', payload[len(payload)-1] == '"'}, false) {
		return calendar.Date{}, calendar.ErrInvalidFormat
	}
	return calendar.ParseDate(string(payload[1 : len(payload)-1]))
}
