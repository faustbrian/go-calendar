package business_test

import (
	"fmt"
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar/v2"
	"github.com/faustbrian/go-calendar/v2/business"
)

// TestRepeatedObservePreservesSourceIdentity retains previously generated
// observance metadata under every later policy while keeping copies independent.
func TestRepeatedObservePreservesSourceIdentity(t *testing.T) {
	original := business.MustHoliday(calendar.MustDate(2026, time.December, 26), "Holiday", map[string]string{"owner": "source"})
	for _, initial := range []business.Observance{business.NextWeekday, business.NearestWeekday} {
		first, err := business.Observe([]business.Holiday{original}, initial)
		if err != nil || len(first) != 2 {
			t.Fatalf("initial Observe = %v", err)
		}
		observed := first[1]
		for _, policy := range []business.Observance{business.NoObservance, business.NextWeekday, business.NearestWeekday} {
			t.Run(fmt.Sprintf("%d/%d", initial, policy), func(t *testing.T) {
				copies, err := business.Observe([]business.Holiday{observed}, policy)
				if err != nil || len(copies) != 1 {
					t.Fatalf("repeated Observe = %v", err)
				}
				copied := copies[0]
				source, ok := copied.SourceDate()
				if !ok || !copied.IsObserved() || source != original.Date() {
					t.Fatalf("prior observation lost: source=%s, observed=%v", source, copied.IsObserved())
				}
				if copied.Date() != observed.Date() || copied.Name() != observed.Name() || copied.Metadata()["owner"] != "source" {
					t.Fatal("effective holiday data changed")
				}
				metadata := copied.Metadata()
				metadata["owner"] = "caller"
				if observed.Metadata()["owner"] != "source" || copied.Metadata()["owner"] != "source" {
					t.Fatal("metadata copy escaped ownership")
				}
				cal, err := business.NewCalendar(business.Config{Revision: "reobserved", Holidays: copies})
				if err != nil {
					t.Fatal(err)
				}
				retained := cal.Holidays(copied.Date())[0]
				retainedSource, retainedOK := retained.SourceDate()
				if !retainedOK || !retained.IsObserved() || retainedSource != source {
					t.Fatal("calendar copy lost observation identity")
				}
			})
		}
	}
}
