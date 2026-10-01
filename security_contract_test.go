package calendar_test

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar/v2"
	"github.com/faustbrian/go-calendar/v2/business"
	calendartz "github.com/faustbrian/go-calendar/v2/timezone"
)

func TestDateJSONEnvelopePreservesGenericSemanticsAndAtomicity(t *testing.T) {
	date := calendar.MustDate(2024, time.February, 29)
	for _, input := range []string{
		`"2024-02-29"` + strings.Repeat(" ", 52),
		`"\u0032\u0030\u0032\u0034\u002d\u0030\u0032\u002d\u0032\u0039"`,
	} {
		var decoded calendar.Date
		if err := decoded.UnmarshalJSON([]byte(input)); err != nil || decoded != date {
			t.Errorf("legitimate bounded JSON decoded=%v error=%v", decoded, err)
		}
		encoded, err := decoded.MarshalJSON()
		if err != nil || string(encoded) != `"2024-02-29"` {
			t.Errorf("canonical roundtrip=%s error=%v", encoded, err)
		}
	}
	retained := calendar.MustDate(2000, time.January, 1)
	decoded := retained
	input := `"2024-02-29"` + strings.Repeat(" ", 53)
	if err := decoded.UnmarshalJSON([]byte(input)); !errors.Is(err, calendar.ErrInvalidFormat) {
		t.Errorf("65-byte JSON error=%v, want invalid format", err)
	}
	if decoded != retained {
		t.Errorf("rejected JSON changed receiver=%v, want %v", decoded, retained)
	}
}

func TestDateTextRejectsInvalidCanonicalDateAtomically(t *testing.T) {
	retained := calendar.MustDate(2000, time.January, 1)
	decoded := retained
	err := decoded.UnmarshalText([]byte("2024-02-30"))
	if !errors.Is(err, calendar.ErrInvalidFormat) || !errors.Is(err, calendar.ErrInvalidDate) {
		t.Fatalf("invalid calendar date classification=%v", err)
	}
	if decoded != retained {
		t.Fatalf("rejected text changed receiver=%v, want %v", decoded, retained)
	}
	if err := decoded.UnmarshalText([]byte("2024-02-29")); err != nil || decoded != calendar.MustDate(2024, time.February, 29) {
		t.Fatalf("valid leap-day text decoded=%v error=%v", decoded, err)
	}
}

func TestCalendarDiagnosticsOmitCallerComponents(t *testing.T) {
	tests := []struct {
		name    string
		call    func() error
		want    error
		private string
	}{
		{"date", func() error { _, err := calendar.NewDate(12345, time.January, 1); return err }, calendar.ErrInvalidDate, "12345"},
		{"year", func() error { _, err := calendar.NewYear(12345); return err }, calendar.ErrInvalidDate, "12345"},
		{"quarter", func() error { _, err := calendar.NewQuarter(2024, 7); return err }, calendar.ErrInvalidDate, "2024"},
		{"semester", func() error { _, err := calendar.NewSemester(2024, 7); return err }, calendar.ErrInvalidDate, "2024"},
		{"week", func() error { _, err := calendar.NewISOWeek(2024, 99); return err }, calendar.ErrInvalidDate, "2024"},
		{"arithmetic", func() error {
			_, err := calendar.MustDate(2024, time.January, 31).AddMonths(1, calendar.Reject)
			return err
		}, calendar.ErrArithmetic, "2024"},
		{"parsed date", func() error { _, err := calendar.ParseDate("2024-02-30"); return err }, calendar.ErrInvalidFormat, "2024"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.call()
			if !errors.Is(err, test.want) {
				t.Fatalf("classification=%v, want %v", err, test.want)
			}
			if strings.Contains(err.Error(), test.private) {
				t.Errorf("default diagnostic retains caller components: %v", err)
			}
		})
	}
	_, err := calendar.ParseDate("2024-02-30")
	if !errors.Is(err, calendar.ErrInvalidDate) {
		t.Errorf("parsed invalid date lost constructor classification: %v", err)
	}
}

func TestDateJSONDiagnosticIsPrivateWithExplicitSyntaxCause(t *testing.T) {
	retained := calendar.MustDate(2000, time.January, 1)
	decoded := retained
	err := decoded.UnmarshalJSON([]byte("#"))
	if !errors.Is(err, calendar.ErrInvalidFormat) {
		t.Fatalf("classification=%v, want invalid format", err)
	}
	if strings.Contains(err.Error(), "#") {
		t.Errorf("default diagnostic retains rejected character: %v", err)
	}
	var syntax *json.SyntaxError
	if !errors.As(err, &syntax) {
		t.Errorf("explicit syntax classification missing: %v", err)
	}
	if decoded != retained {
		t.Errorf("syntax rejection changed receiver: %v", decoded)
	}
	err = decoded.UnmarshalJSON([]byte("#" + strings.Repeat(" ", 64)))
	syntax = nil
	if !errors.Is(err, calendar.ErrInvalidFormat) || errors.As(err, &syntax) {
		t.Errorf("oversized JSON must reject before syntax decoding: %v", err)
	}
	err = decoded.UnmarshalJSON([]byte("123"))
	var valueType *json.UnmarshalTypeError
	if !errors.Is(err, calendar.ErrInvalidFormat) || !errors.As(err, &valueType) {
		t.Errorf("explicit type classification missing: %v", err)
	}
	if decoded != retained {
		t.Errorf("type or size rejection changed receiver: %v", decoded)
	}
}

func TestBusinessWeekendAdmissionBound(t *testing.T) {
	sunday := calendar.MustDate(2024, time.January, 7)
	monday := calendar.MustDate(2024, time.January, 8)
	for _, weekdays := range [][]time.Weekday{
		{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday},
		{time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday},
	} {
		cal, err := business.NewCalendar(business.Config{Revision: "ordinary", Weekends: weekdays})
		if err != nil || !cal.IsValid() || cal.IsBusinessDay(sunday) {
			t.Fatalf("seven weekdays rejected or Sunday left open: %v", err)
		}
		if got, want := cal.IsBusinessDay(monday), weekdays[1] == time.Sunday; got != want {
			t.Errorf("seven weekdays Monday open=%v, want %v", got, want)
		}
	}
	oversized := []time.Weekday{time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday, time.Sunday}
	cal, err := business.NewCalendar(business.Config{Revision: "ordinary", Weekends: oversized})
	if !errors.Is(err, business.ErrResourceLimit) || cal.IsValid() || cal.Revision() != "" {
		t.Errorf("eight weekdays result valid=%v revision=%q error=%v, want zero calendar and resource limit", cal.IsValid(), cal.Revision(), err)
	}
	if _, err := business.NewCalendar(business.Config{Weekends: oversized}); !errors.Is(err, business.ErrInvalidCalendar) {
		t.Errorf("revision-first precedence lost: %v", err)
	}
	if _, err := business.NewCalendar(business.Config{Revision: "ordinary", Weekends: []time.Weekday{8}}); !errors.Is(err, business.ErrInvalidCalendar) {
		t.Errorf("invalid weekday classification lost: %v", err)
	}
}

func TestBoundedStringAdmissionPreservesClassification(t *testing.T) {
	if location, err := calendartz.LoadLocation(strings.Repeat("x", calendartz.MaxZoneNameBytes+1)); location != nil || !errors.Is(err, calendartz.ErrInvalidZone) {
		t.Errorf("oversized zone classification=%v location=%v", err, location)
	}
	if _, err := business.NewCalendar(business.Config{Revision: strings.Repeat("r", 129)}); !errors.Is(err, business.ErrInvalidCalendar) {
		t.Errorf("oversized revision classification=%v", err)
	}
	date := calendar.MustDate(2024, time.January, 1)
	for _, metadata := range []map[string]string{
		{strings.Repeat("k", 129): "ordinary"},
		{"ordinary": strings.Repeat("v", 1025)},
	} {
		if _, err := business.NewHoliday(date, "ordinary", metadata); !errors.Is(err, business.ErrInvalidHoliday) {
			t.Errorf("oversized metadata classification=%v", err)
		}
	}
}
