package calendarwire_test

import (
	"testing"

	calendarwire "github.com/faustbrian/go-calendar/v2/adapters/wire"
)

// FuzzDecodeDate checks successful bounded wire decoding re-encodes to identical bytes.
func FuzzDecodeDate(f *testing.F) {
	for _, seed := range [][]byte{[]byte(`"2024-02-29"`), []byte("null"), []byte(`"bad"`), {0xff}} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, payload []byte) {
		date, err := calendarwire.DecodeDate(payload)
		if err != nil {
			return
		}
		encoded, err := calendarwire.EncodeDate(date)
		if err != nil || string(encoded) != string(payload) {
			t.Fatalf("wire round trip = %q, %v", encoded, err)
		}
	})
}
