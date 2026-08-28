# calendar

[![CI](https://github.com/faustbrian/go-calendar/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/faustbrian/go-calendar/actions/workflows/ci.yml)
[![CodeQL](https://img.shields.io/badge/CodeQL-required-blue)](https://github.com/faustbrian/go-calendar/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Mutation](https://img.shields.io/badge/mutation-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Documentation](https://img.shields.io/badge/docs-checked_in_CI-blue)](docs/)
[![Go Reference](https://pkg.go.dev/badge/github.com/faustbrian/go-calendar.svg)](https://pkg.go.dev/github.com/faustbrian/go-calendar)
[![Release](https://img.shields.io/github/v/release/faustbrian/go-calendar?sort=semver)](https://github.com/faustbrian/go-calendar/releases)
[![Go](https://img.shields.io/badge/go-1.26.6-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`calendar` provides immutable civil dates, Gregorian calendar arithmetic,
typed calendar periods, explicit DST conversion, and bounded business
calendars for Go 1.26.6 and later.

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

## Boundaries

- `clock` owns current time; [calendarclock](calendarclock) accepts its
  narrow `Now() time.Time` capability.
- `temporal` owns interval and set algebra; [calendartemporal](calendartemporal)
  supplies explicit instant boundaries and bounded date sequences.
- This project does not provide clocks, timers, cron, scheduling, natural
  language parsing, opening hours, or a global holiday database.

## Quality

`golib check --all` reproduces every blocking local gate and reports advisory
NilAway findings. Production packages maintain meaningful 100.0%
statement coverage; `calendartest` is test-support code and is excluded from
that denominator. See [verification](docs/verification.md).

## License

MIT. See [LICENSE](LICENSE).
