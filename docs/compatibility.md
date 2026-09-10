# Compatibility

The minimum toolchain is Go 1.27.0, the latest stable release at implementation
time. The module supports standard-library IANA behavior on Linux, macOS, and
Windows. Timezone results follow the installed or embedded tzdata snapshot.

PostgreSQL integration targets maintained versions 14 through 18. pgx v5.10.0
is pinned. Public API drift is checked against `api/baseline.txt`; deliberate
breaking changes require a major version after v1.

The six canonical `adapters/*` packages are additive root-module APIs. Their
former top-level paths remain supported for the longer of 180 days after the
canonical successor becomes public and two later stable root-module minor
releases containing both paths. Removal additionally requires owned-consumer
migration, external-consumer evidence, and a future authorized major release.
