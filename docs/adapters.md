# Adapters

Canonical adapters use target-oriented paths:

| Capability | Canonical import | Deprecated compatibility import |
| --- | --- | --- |
| clock | `github.com/faustbrian/go-calendar/v2/adapters/clock` | `github.com/faustbrian/go-calendar/v2/calendarclock` |
| config | `github.com/faustbrian/go-calendar/v2/adapters/config` | `github.com/faustbrian/go-calendar/v2/calendarconfig` |
| PostgreSQL | `github.com/faustbrian/go-calendar/v2/adapters/postgres` | `github.com/faustbrian/go-calendar/v2/postgres` |
| temporal | `github.com/faustbrian/go-calendar/v2/adapters/temporal` | `github.com/faustbrian/go-calendar/v2/calendartemporal` |
| validation | `github.com/faustbrian/go-calendar/v2/adapters/validation` | `github.com/faustbrian/go-calendar/v2/calendarvalidation` |
| wire | `github.com/faustbrian/go-calendar/v2/adapters/wire` | `github.com/faustbrian/go-calendar/v2/calendarwire` |

`calendarconfig.Date` from `adapters/config` accepts only a string and implements both config's
`UnmarshalConfigValue(any)` seam and standard text unmarshalling. Null and
numeric coercion are rejected.

`calendarvalidation.ValidDate` and `DateRange` from `adapters/validation`
return dependency-neutral rules.
Wrap a rule in the current `validation.ValidatorFunc` and translate its safe
sentinel error into the application's violation code. This avoids a circular or
unpublished module dependency.

`calendarwire.EncodeDate` and `DecodeDate` from `adapters/wire` provide the
bounded wire-version-1 JSON contract, independent of the module's v2 identity.
The root `Date` also composes directly with wire formats that honor standard
text or JSON encoding interfaces.

`adapters/clock` borrows the supplied clock and location for one call.
`adapters/temporal` returns immutable instant boundaries and bounded date
sequences. None of these adapters starts background work or owns shutdown.

See [adapter migration](migration.md) for the additive import-path change.
