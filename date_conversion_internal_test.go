package calendar

import (
	"testing"
	"time"
)

// asTime is private; public Weekday rejects invalid dates before calling it.
// Keep this defensive conversion contract without inventing a public route.
func TestPrivateDateConversionRefusesInvalidCivilValues(t *testing.T) {
	for _, date := range []Date{
		{},
		{year: 2023, month: 2, day: 29},
		{year: 2024, month: 13, day: 1},
		{year: 10000, month: 1, day: 1},
	} {
		if date.IsValid() {
			t.Fatal("invalid control unexpectedly valid")
		}
		if converted := date.asTime(); !converted.IsZero() {
			t.Fatalf("invalid private date converted to %v", converted)
		}
	}
	for _, date := range []Date{
		MustDate(MinYear, time.January, 1),
		MustDate(2024, time.February, 29),
		MustDate(MaxYear, time.December, 31),
	} {
		converted := date.asTime()
		expected := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
		if !converted.Equal(expected) || converted.Location() != time.UTC {
			t.Fatalf("%s converted to %v; want exact UTC midnight %v", date, converted, expected)
		}
	}
}
