//nolint:staticcheck // This compatibility test intentionally exercises deprecated facades.
package calendar_test

//lint:file-ignore SA1019 This compatibility test intentionally exercises deprecated facades.

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar/v2"
	calendarclock "github.com/faustbrian/go-calendar/v2/adapters/clock"
	calendarconfig "github.com/faustbrian/go-calendar/v2/adapters/config"
	calendarpostgres "github.com/faustbrian/go-calendar/v2/adapters/postgres"
	calendartemporal "github.com/faustbrian/go-calendar/v2/adapters/temporal"
	calendarvalidation "github.com/faustbrian/go-calendar/v2/adapters/validation"
	calendarwire "github.com/faustbrian/go-calendar/v2/adapters/wire"
	legacyclock "github.com/faustbrian/go-calendar/v2/calendarclock"
	legacyconfig "github.com/faustbrian/go-calendar/v2/calendarconfig"
	legacytemporal "github.com/faustbrian/go-calendar/v2/calendartemporal"
	legacyvalidation "github.com/faustbrian/go-calendar/v2/calendarvalidation"
	legacywire "github.com/faustbrian/go-calendar/v2/calendarwire"
	legacypostgres "github.com/faustbrian/go-calendar/v2/postgres"
)

type successorFixedClock struct{ now time.Time }

func (c successorFixedClock) Now() time.Time { return c.now }

func TestSuccessorAdaptersExposeFrozenContracts(t *testing.T) {
	t.Parallel()

	date := calendar.MustDate(2024, time.February, 29)
	clock := successorFixedClock{now: time.Date(2024, time.February, 29, 23, 0, 0, 0, time.UTC)}
	if today, err := calendarclock.Today(clock, time.UTC); err != nil || !today.Equal(date) {
		t.Fatalf("Today() = %v, %v", today, err)
	}

	configDate := calendarconfig.NewDate(date)
	if got := configDate.CalendarDate(); !got.Equal(date) {
		t.Fatalf("config date = %v", got)
	}

	sequence, err := calendartemporal.Sequence(date, date, 1)
	if err != nil || len(sequence) != 1 || !sequence[0].Equal(date) {
		t.Fatalf("Sequence() = %v, %v", sequence, err)
	}

	if err := calendarvalidation.ValidDate()(date); err != nil {
		t.Fatalf("ValidDate() = %v", err)
	}

	payload, err := calendarwire.EncodeDate(date)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := calendarwire.DecodeDate(payload)
	if err != nil || !decoded.Equal(date) {
		t.Fatalf("DecodeDate() = %v, %v", decoded, err)
	}

	value, err := calendarpostgres.NewDate(date).Value()
	if err != nil || value != "2024-02-29" {
		t.Fatalf("Value() = %v, %v", value, err)
	}
}

func TestLegacyAdaptersRemainDistinctCompatibleFacades(t *testing.T) {
	t.Parallel()

	if reflect.TypeOf(legacyconfig.Date{}).PkgPath() != "github.com/faustbrian/go-calendar/v2/calendarconfig" {
		t.Fatal("legacy config Date lost its package identity")
	}
	if reflect.TypeOf(legacytemporal.InstantRange{}).PkgPath() != "github.com/faustbrian/go-calendar/v2/calendartemporal" {
		t.Fatal("legacy temporal InstantRange lost its package identity")
	}
	if reflect.TypeOf(legacyvalidation.Rule(nil)).PkgPath() != "github.com/faustbrian/go-calendar/v2/calendarvalidation" {
		t.Fatal("legacy validation Rule lost its package identity")
	}
	if reflect.TypeOf(legacypostgres.Date{}).PkgPath() != "github.com/faustbrian/go-calendar/v2/postgres" {
		t.Fatal("legacy postgres Date lost its package identity")
	}
	if reflect.TypeOf(legacypostgres.InfinityKind(0)).PkgPath() != "github.com/faustbrian/go-calendar/v2/postgres" {
		t.Fatal("legacy postgres InfinityKind lost its package identity")
	}

	if !errors.Is(legacyclock.ErrClockRequired, calendarclock.ErrClockRequired) ||
		!errors.Is(legacyvalidation.ErrInvalidDate, calendarvalidation.ErrInvalidDate) ||
		!errors.Is(legacywire.ErrSizeLimit, calendarwire.ErrSizeLimit) ||
		!errors.Is(legacypostgres.ErrNull, calendarpostgres.ErrNull) {
		t.Fatal("legacy and successor sentinel identities differ")
	}
}

func TestDateTextVariantsPreserveAtomicCanonicalAdmission(t *testing.T) {
	retained := calendar.MustDate(2000, time.January, 1)
	root := retained
	active := calendarconfig.NewDate(retained)
	legacy := legacyconfig.NewDate(retained)
	for _, target := range []interface{ UnmarshalText([]byte) error }{&root, &active, &legacy} {
		if err := target.UnmarshalText([]byte("2024-02-29 ")); !errors.Is(err, calendar.ErrInvalidFormat) {
			t.Errorf("overlength text classification=%v", err)
		}
	}
	if root != retained || active.CalendarDate() != retained || legacy.CalendarDate() != retained {
		t.Fatal("rejected text changed a receiver")
	}
	for _, target := range []interface{ UnmarshalText([]byte) error }{&root, &active, &legacy} {
		if err := target.UnmarshalText([]byte("2024-02-29")); err != nil {
			t.Errorf("canonical text rejection: %v", err)
		}
	}
	date := calendar.MustDate(2024, time.February, 29)
	if root != date || active.CalendarDate() != date || legacy.CalendarDate() != date {
		t.Fatal("canonical text did not reach every receiver")
	}
}

func TestUnsupportedValueDiagnosticsArePrivateAcrossAdapterVariants(t *testing.T) {
	retained := calendar.MustDate(2000, time.January, 1)
	activeConfig := calendarconfig.NewDate(retained)
	legacyConfig := legacyconfig.NewDate(retained)
	activePostgres := calendarpostgres.NewDate(retained)
	legacyPostgres := legacypostgres.NewDate(retained)
	input := struct{ RequestField string }{RequestField: "ordinary"}
	tests := []struct {
		name   string
		decode func(any) error
		value  func() calendar.Date
	}{
		{"active config", activeConfig.UnmarshalConfigValue, func() calendar.Date { return activeConfig.CalendarDate() }},
		{"legacy config", legacyConfig.UnmarshalConfigValue, func() calendar.Date { return legacyConfig.CalendarDate() }},
		{"active postgres", activePostgres.Scan, func() calendar.Date { return activePostgres.CalendarDate() }},
		{"legacy postgres", legacyPostgres.Scan, func() calendar.Date { return legacyPostgres.CalendarDate() }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.decode(input)
			if err == nil {
				t.Fatal("unsupported value accepted")
			}
			if strings.Contains(err.Error(), "RequestField") {
				t.Errorf("default diagnostic retains caller type details: %v", err)
			}
			if errors.Is(err, calendar.ErrInvalidFormat) || errors.Is(err, calendarpostgres.ErrNull) || errors.Is(err, calendarpostgres.ErrInfinity) {
				t.Errorf("unsupported value reclassified as format, NULL, or infinity: %v", err)
			}
			if test.value() != retained {
				t.Errorf("unsupported value changed receiver: %v", test.value())
			}
		})
	}
}
