package calendar_test

import (
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar/v2"
	"github.com/faustbrian/go-calendar/v2/business"
)

// TestInvalidDateWeekdayUsesDocumentedZero preserves the original Sunday
// fallback without treating an invalid civil date as a valid business day.
func TestInvalidDateWeekdayUsesDocumentedZero(t *testing.T) {
	zero := calendar.Date{}
	if got := zero.Weekday(); got != time.Sunday {
		t.Fatalf("invalid Weekday = %v, want Sunday", got)
	}
	for day := 1; day <= 7; day++ {
		date := calendar.MustDate(1, time.January, day)
		if got, want := date.Weekday(), time.Weekday(day%7); got != want {
			t.Fatalf("%s Weekday = %v, want %v", date, got, want)
		}
	}
	policy, _ := calendar.NewWeekPolicy(time.Sunday)
	if policy.StartOfWeek(zero).IsValid() || zero.StartOfISOWeek().IsValid() {
		t.Fatal("invalid date acquired a week boundary")
	}
	cal, err := business.NewCalendar(business.Config{Revision: "invalid-weekday", Weekends: []time.Weekday{time.Sunday}})
	if err != nil {
		t.Fatal(err)
	}
	if cal.IsWeekend(zero) || cal.IsBusinessDay(zero) {
		t.Fatal("invalid date admitted by business calendar")
	}
}

// TestBoundaryISOWeekRetainsCivilRange characterizes mathematical week identity
// separately from representability of its Sunday in the finite Date domain.
func TestBoundaryISOWeekRetainsCivilRange(t *testing.T) {
	week, err := calendar.NewISOWeek(9999, 52)
	if err != nil || !week.IsValid() {
		t.Fatalf("boundary ISOWeek = %v, %v", week, err)
	}
	if got := week.FirstDate().String(); got != "9999-12-27" {
		t.Fatalf("first date = %q", got)
	}
	if week.LastDate().IsValid() {
		t.Fatalf("Sunday escaped civil range: %s", week.LastDate())
	}
	maximum := calendar.MustDate(9999, time.December, 31)
	if year, number := maximum.ISOWeek(); year != week.Year() || number != week.Week() || !week.Contains(maximum) {
		t.Fatal("representable final date lost ISO week identity")
	}
	if maximum.EndOfISOWeek().IsValid() || maximum.EndOfISOYear().IsValid() {
		t.Fatal("unrepresentable week boundary was clipped or admitted")
	}
	if week.Contains(calendar.Date{}) {
		t.Fatal("invalid date admitted by boundary week")
	}
	ordinary, err := calendar.NewISOWeek(9999, 51)
	if err != nil || ordinary.LastDate().String() != "9999-12-26" {
		t.Fatalf("previous week Sunday = %s, %v", ordinary.LastDate(), err)
	}
}
