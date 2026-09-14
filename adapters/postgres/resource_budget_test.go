package calendarpostgres_test

import (
	"errors"
	"testing"
	"time"

	calendar "github.com/faustbrian/go-calendar"
	calendarpg "github.com/faustbrian/go-calendar/adapters/postgres"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	budgetPGXBytes []byte
	budgetPGXErr   error
)

func TestPGXBinaryEncodingAllocationBudget(t *testing.T) {
	codec := pgtype.NewMap()
	value := calendarpg.NewDate(calendar.MustDate(2024, time.February, 29))
	if allocations := testing.AllocsPerRun(1_000, func() {
		budgetPGXBytes, budgetPGXErr = codec.Encode(
			pgtype.DateOID, pgtype.BinaryFormatCode, value, nil,
		)
	}); allocations > 2 {
		t.Fatalf("pgx binary encode allocations = %.0f, budget 2", allocations)
	}
}

func TestOversizedSQLByteInputAllocationBudget(t *testing.T) {
	payload := make([]byte, 1<<20)
	for name, scanner := range map[string]interface{ Scan(any) error }{
		"date":     &calendarpg.Date{},
		"infinity": &calendarpg.InfinityDate{},
	} {
		t.Run(name, func(t *testing.T) {
			result := testing.Benchmark(func(b *testing.B) {
				for range b.N {
					budgetPGXErr = scanner.Scan(payload)
				}
			})
			if bytes := result.AllocedBytesPerOp(); bytes > 1<<10 {
				t.Fatalf("oversized SQL byte scan allocated %d bytes, budget %d", bytes, 1<<10)
			}
			if !errors.Is(budgetPGXErr, calendar.ErrInvalidFormat) {
				t.Fatalf("oversized SQL byte scan error = %v", budgetPGXErr)
			}
		})
	}
}
