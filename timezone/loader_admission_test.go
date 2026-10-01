package timezone

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestLoadLocationAdmitsInclusiveNameBeforeLoader(t *testing.T) {
	name := strings.Repeat("a", 255)
	want := time.FixedZone("provided", 3600)
	calls := 0
	location, err := loadLocation(name, func(got string) (*time.Location, error) {
		calls++
		if got != name {
			t.Fatalf("loader name = %q, want unchanged input", got)
		}
		return want, nil
	})
	if calls != 1 || location != want || err != nil {
		t.Fatalf("inclusive admission: calls=%d location=%v error=%v", calls, location, err)
	}
}

func TestLoadLocationRejectsBeforeLoader(t *testing.T) {
	for _, test := range []struct{ name, input string }{
		{"empty", ""},
		{"over byte limit", strings.Repeat("a", 256)},
		{"invalid UTF8", "before\xffafter"},
		{"absolute path", "/zone"},
		{"backslash", "zone\\name"},
		{"empty segment", "zone//name"},
		{"dot segment", "zone/./name"},
		{"parent segment", "zone/../name"},
	} {
		t.Run(test.name, func(t *testing.T) {
			calls := 0
			location, err := loadLocation(test.input, func(string) (*time.Location, error) {
				calls++
				return time.UTC, nil
			})
			if calls != 0 || location != nil || !errors.Is(err, ErrInvalidZone) {
				t.Fatalf("rejected admission: calls=%d location=%v error=%v", calls, location, err)
			}
		})
	}
}

func TestLoadLocationKeepsLoaderFailurePrivate(t *testing.T) {
	private := errors.New("application-private-loader-detail")
	calls := 0
	location, err := loadLocation("zone", func(name string) (*time.Location, error) {
		calls++
		if name != "zone" {
			t.Fatalf("loader name = %q", name)
		}
		return time.UTC, private
	})
	if calls != 1 || location != nil || !errors.Is(err, ErrInvalidZone) || errors.Is(err, private) {
		t.Fatalf("loader failure: calls=%d location=%v error=%v", calls, location, err)
	}
	if strings.Contains(err.Error(), "application-private-loader-detail") {
		t.Fatal("loader failure exposed private detail")
	}
}
