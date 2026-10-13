# PostgreSQL

[Documentation map](index.md) · [Repository README](../README.md)
[Next: API reference](api.md)

Import `github.com/faustbrian/go-calendar/v2/adapters/postgres` and use
`calendarpostgres.NewDate` for a finite non-null PostgreSQL `date`. The adapter
implements `database/sql.Scanner`, `driver.Valuer`, `pgtype.DateScanner`, and
`pgtype.DateValuer`. SQL `NULL` and infinity are errors.

Use `calendarpostgres.InfinityDate` only when the schema intentionally admits
`-infinity` or `infinity`. Those sentinels are never coerced to ordinary civil
dates; InfinityDate still rejects NULL. Scans preserve the receiver on error
and require exclusive mutation ownership. Finite time.Time/pgx inputs use civil
components in their supplied location, without automatic UTC conversion.
Native pgx text/binary codecs and a live PostgreSQL round trip are tested.

`golib check --module .` starts a task-owned PostgreSQL fixture and runs the
tagged integration suite. CI runs the same contract against the pinned service
identity.

The former `github.com/faustbrian/go-calendar/v2/postgres` import remains a
deprecated compatibility path. See [adapter migration](migration.md).
