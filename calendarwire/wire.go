// Package calendarwire provides bounded canonical wire helpers for Date.
// EncodeDate returns owned bytes. DecodeDate borrows input only during its call
// and accepts exactly the 12-byte unescaped JSON date string, not the broader
// generic Date JSON grammar. Version is external contract metadata, not an
// envelope byte; upstream payload acquisition/decompression remains caller-owned.
//
// Deprecated: use github.com/faustbrian/go-calendar/v2/adapters/wire. This package
// remains supported for the longer of 180 days after successor public
// availability and two subsequently published stable root-module minor
// releases.
package calendarwire

import (
	"cmp"
	"slices"

	calendar "github.com/faustbrian/go-calendar/v2"
	adapter "github.com/faustbrian/go-calendar/v2/adapters/wire"
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
