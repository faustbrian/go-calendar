package calendarwire_test

import (
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar"
	"github.com/faustbrian/go-calendar/calendarwire"
)

var (
	budgetWireBytes []byte
	budgetWireDate  calendar.Date
	budgetWireErr   error
)

func TestWireOutputAndAllocationBudgets(t *testing.T) {
	date := calendar.MustDate(2024, time.February, 29)
	maximumEncodeAllocations := float64(5 + raceAllocationSlack)
	if allocations := testing.AllocsPerRun(1_000, func() {
		budgetWireBytes, budgetWireErr = calendarwire.EncodeDate(date)
	}); allocations > maximumEncodeAllocations {
		t.Fatalf("wire encode allocations = %.0f, budget %.0f", allocations, maximumEncodeAllocations)
	}
	if budgetWireErr != nil || len(budgetWireBytes) > calendarwire.MaxBytes {
		t.Fatalf("wire output = %d bytes, %v", len(budgetWireBytes), budgetWireErr)
	}
	payload := []byte(`"2024-02-29"`)
	maximumDecodeAllocations := float64(4 + raceAllocationSlack)
	if allocations := testing.AllocsPerRun(1_000, func() {
		budgetWireDate, budgetWireErr = calendarwire.DecodeDate(payload)
	}); allocations > maximumDecodeAllocations {
		t.Fatalf("wire decode allocations = %.0f, budget %.0f", allocations, maximumDecodeAllocations)
	}
}
