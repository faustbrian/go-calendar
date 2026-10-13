# Versioning and persisted meaning

[Documentation map](index.md) · [Repository README](../README.md)
[Next: Five-minute guide](quickstart.md)

Calendar APIs follow semantic versioning after v1. Business calendars carry an
application-supplied revision independently from the Go module version.

Persist civil values in canonical form. Persist business revision and dataset
provenance with derived decisions. Persist IANA zone identity, resolution
policy, and tzdata/application version when local-to-instant replay matters.
Never reinterpret an old decision silently under a new calendar revision.

## v2.0.1 patch candidate

The candidate keeps the `github.com/faustbrian/go-calendar/v2` module and import
paths and the Go 1.27.0 public minimum. Once published, update the module to
`v2.0.1`; existing v2 imports need no migration. Preparation is not publication.

Invalid Date weekday queries regain the documented Sunday fallback; callers
must still check `IsValid` before using zero values as actual dates. Holiday
observance copies now retain their original source identity, but reprocessing
weekend holidays can still append duplicates. The civil range remains years
1–9999: a valid `9999-W52` still has an unrepresentable final Sunday, so check
period endpoint validity before constructing complete date ranges.
