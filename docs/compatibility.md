# Compatibility

The minimum toolchain is Go 1.27.0, the latest stable release at implementation
time. The module supports standard-library IANA behavior on Linux, macOS, and
Windows. Timezone results follow the installed or embedded tzdata snapshot.

PostgreSQL integration targets maintained versions 14 through 18. pgx v5.10.0
is pinned. Public API drift is checked against `api/calendar-v2.txt`; deliberate
breaking changes require a major version after v1.

Main contains v2.0.0 with the official `/v2` module suffix. The unchanged
`api/baseline.txt` records the historical v1 baseline, not the active v2 API.
Published v1.1.0 remains separate from the v2 source. Public tags and releases
establish v2 publication status. See
[v2 migration](v2-migration.md) for decoding, diagnostics, and type-identity
changes.

The six canonical `adapters/*` packages are additive root-module APIs. Their
former top-level paths remain supported for the longer of 180 days after the
canonical successor becomes public and two later stable root-module minor
releases containing both paths. Removal additionally requires owned-consumer
migration, external-consumer evidence, and a future authorized major release.
