# Adapter migration

The target-oriented adapter paths are additive. Existing programs may migrate
one import at a time without changing constructors, methods, constants,
sentinel checks, serialization, ownership, or concurrency behavior.

| Replace | With |
| --- | --- |
| `go-calendar/calendarclock` | `go-calendar/adapters/clock` |
| `go-calendar/calendarconfig` | `go-calendar/adapters/config` |
| `go-calendar/postgres` | `go-calendar/adapters/postgres` |
| `go-calendar/calendartemporal` | `go-calendar/adapters/temporal` |
| `go-calendar/calendarvalidation` | `go-calendar/adapters/validation` |
| `go-calendar/calendarwire` | `go-calendar/adapters/wire` |

The package identifiers remain `calendarclock`, `calendarconfig`,
`calendartemporal`, `calendarvalidation`, and `calendarwire`. The canonical
PostgreSQL identifier is `calendarpostgres`; consumers that previously relied
on the default `postgres` identifier must update that qualifier or retain an
explicit import alias.

The old packages remain importable and preserve their released named-type and
sentinel identities. They are deprecated compatibility paths; new code should
not introduce them. The wire compatibility path retains its small local codec
implementation because an additional package-call boundary exceeds its
published allocation budget.
