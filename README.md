# calendar

[![CI](https://github.com/faustbrian/go-calendar/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/faustbrian/go-calendar/actions/workflows/ci.yml)
[![Documentation](https://img.shields.io/badge/docs-adoption_and_API-blue)](docs/index.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/faustbrian/go-calendar/v2.svg)](https://pkg.go.dev/github.com/faustbrian/go-calendar/v2)
[![Release](https://img.shields.io/github/v/release/faustbrian/go-calendar?sort=semver)](https://github.com/faustbrian/go-calendar/releases)
[![Go](https://img.shields.io/badge/go-1.27.0-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`calendar` provides immutable civil dates, Gregorian calendar arithmetic,
typed calendar periods, explicit DST conversion, and bounded business
calendars for Go 1.27.0 and later.

The published v2.0.0 module uses `github.com/faustbrian/go-calendar/v2`. Public tags
and GitHub releases establish publication status; source documentation alone
does not. V1.1.0 remains available under the original module path.
See the [v2 migration guide](docs/v2-migration.md) before changing
an existing application or public adapter.

A `calendar.Date` is a day on a calendar. It is not a `time.Time`, has no
timezone, and cannot accidentally be used as an elapsed duration.

```go
date := calendar.MustDate(2024, time.January, 31)
next, err := date.AddMonths(1, calendar.Clamp) // 2024-02-29
```

Conversion to an instant always names both an IANA location and a resolution
policy:

```go
location, _ := calendartz.LoadLocation("America/New_York")
local := calendartz.MustLocalDateTime(date, 9, 0, 0, 0)
instant, err := calendartz.Resolve(local, location, calendartz.Reject)
```

Business calendars are application-supplied, immutable, revisioned, and
search-bounded:

```go
cal, _ := business.NewCalendar(business.Config{
    Revision: "company-2024.1",
    Weekends: []time.Weekday{time.Saturday, time.Sunday},
})
due, err := cal.AddBusinessDays(date, 5, 14)
```

Start with the [five-minute guide](docs/quickstart.md), then use the
[API guide](docs/api.md), [timezone guide](docs/timezone.md), and
[business-calendar guide](docs/business.md).

## Adoption entry points

The root module is `github.com/faustbrian/go-calendar/v2` (stable).
It requires Go 1.27.0. Pin the published `v2.0.0` root tag
with `go get github.com/faustbrian/go-calendar/v2@v2.0.0`.
Root subpackages share this module; they need no separate installation.

Start with the [documentation map](docs/index.md) for package selection,
examples, operational constraints, and contributor references.

## Boundaries

- `clock` owns current time; [adapters/clock](adapters/clock) accepts its
  narrow `Now() time.Time` capability.
- `temporal` owns interval and set algebra; [adapters/temporal](adapters/temporal)
  supplies explicit instant boundaries and bounded date sequences.
- Canonical configuration, validation, wire, and PostgreSQL integrations live
  under [`adapters/`](docs/adapters.md). The original top-level adapter paths
  remain supported compatibility facades during the documented migration
  interval.
- Caller-provided clocks, locations, policies, and byte slices are borrowed
  only for the call and are not retained. Business configuration maps and
  slices are defensively copied.
- This project does not provide clocks, timers, cron, scheduling, natural
  language parsing, opening hours, or a global holiday database.

## Quality

Start standalone changes with `go test ./...`. With the qualified public
verifier, `make check` and `make ci` select bounded local checks;
`make check-full` preserves explicit full assurance. The manifest's PostgreSQL
service and integration tag make the Make route distinct from untagged Go tests.
See [contributor verification](CONTRIBUTING.md#verification) for focused fuzzing,
optional local mutation, service and release boundaries, and
[verification](docs/verification.md) for the retained behavioral evidence.

For ecosystem-wide selection and ownership guidance, see the versioned
[Golib ecosystem index](https://github.com/faustbrian/go-library-tools/blob/v1.8.4/docs/ecosystem/README.md)
and its [Domain utilities family](https://github.com/faustbrian/go-library-tools/blob/v1.8.4/docs/ecosystem/design-language.md#package-families-and-selection).

## License

MIT. See [LICENSE](LICENSE).
