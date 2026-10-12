package calendar_test

import (
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar/v2"
)

// BenchmarkParseDate measures one canonical leap-day parse per iteration.
func BenchmarkParseDate(b *testing.B) {
	for b.Loop() {
		_, _ = calendar.ParseDate("2024-02-29")
	}
}

// BenchmarkAddMonths measures clamped January-end movement with setup outside b.Loop.
func BenchmarkAddMonths(b *testing.B) {
	date := calendar.MustDate(2024, time.January, 31)
	for b.Loop() {
		_, _ = date.AddMonths(1, calendar.Clamp)
	}
}

// BenchmarkISOWeek measures one boundary-date ISO query with setup outside b.Loop.
func BenchmarkISOWeek(b *testing.B) {
	date := calendar.MustDate(2024, time.December, 31)
	for b.Loop() {
		_, _ = date.ISOWeek()
	}
}
