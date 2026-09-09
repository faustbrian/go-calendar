//nolint:staticcheck // This compatibility test intentionally exercises deprecated facades.
package calendar_test

//lint:file-ignore SA1019 This compatibility test intentionally exercises deprecated facades.

import (
	"errors"
	"reflect"
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar"
	calendarclock "github.com/faustbrian/go-calendar/adapters/clock"
	calendarconfig "github.com/faustbrian/go-calendar/adapters/config"
	calendarpostgres "github.com/faustbrian/go-calendar/adapters/postgres"
	calendartemporal "github.com/faustbrian/go-calendar/adapters/temporal"
	calendarvalidation "github.com/faustbrian/go-calendar/adapters/validation"
	calendarwire "github.com/faustbrian/go-calendar/adapters/wire"
	legacyclock "github.com/faustbrian/go-calendar/calendarclock"
	legacyconfig "github.com/faustbrian/go-calendar/calendarconfig"
	legacytemporal "github.com/faustbrian/go-calendar/calendartemporal"
	legacyvalidation "github.com/faustbrian/go-calendar/calendarvalidation"
	legacywire "github.com/faustbrian/go-calendar/calendarwire"
	legacypostgres "github.com/faustbrian/go-calendar/postgres"
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

	if reflect.TypeOf(legacyconfig.Date{}).PkgPath() != "github.com/faustbrian/go-calendar/calendarconfig" {
		t.Fatal("legacy config Date lost its package identity")
	}
	if reflect.TypeOf(legacytemporal.InstantRange{}).PkgPath() != "github.com/faustbrian/go-calendar/calendartemporal" {
		t.Fatal("legacy temporal InstantRange lost its package identity")
	}
	if reflect.TypeOf(legacyvalidation.Rule(nil)).PkgPath() != "github.com/faustbrian/go-calendar/calendarvalidation" {
		t.Fatal("legacy validation Rule lost its package identity")
	}
	if reflect.TypeOf(legacypostgres.Date{}).PkgPath() != "github.com/faustbrian/go-calendar/postgres" {
		t.Fatal("legacy postgres Date lost its package identity")
	}
	if reflect.TypeOf(legacypostgres.InfinityKind(0)).PkgPath() != "github.com/faustbrian/go-calendar/postgres" {
		t.Fatal("legacy postgres InfinityKind lost its package identity")
	}

	if !errors.Is(legacyclock.ErrClockRequired, calendarclock.ErrClockRequired) ||
		!errors.Is(legacyvalidation.ErrInvalidDate, calendarvalidation.ErrInvalidDate) ||
		!errors.Is(legacywire.ErrSizeLimit, calendarwire.ErrSizeLimit) ||
		!errors.Is(legacypostgres.ErrNull, calendarpostgres.ErrNull) {
		t.Fatal("legacy and successor sentinel identities differ")
	}
}
