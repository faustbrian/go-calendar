# Changelog

All notable changes follow Keep a Changelog. The project uses semantic
versioning.

## Unreleased

### Changed

- Track the pinned documentation-tool lockfile so clean CI checkouts install
  the exact validated cspell dependency.

- Reconcile standalone dependency checksums against deterministic current
  module archives so CI, local verification, and release consumers resolve
  identical content.

- Harden standalone documentation validation with deterministic spelling and
  link checks, package-specific documentation gates, and repository-local
  contributor guidance.

## 1.0.0 - 2026-08-25

### Documentation

- Link the package README to the repository-wide Golib documentation portal.

### Changed

- Publish the module from its standalone `github.com/faustbrian/go-calendar` identity while preserving its documented API and behavior.
- Replaced host Ruby documentation validation with a self-contained Go link
  checker that uses the module's declared toolchain.

### Added

- Immutable bounded `Date` and typed calendar periods.
- Explicit clamp, reject, and overflow arithmetic policies.
- DST gap/fold detection and bounded IANA loading.
- Immutable revisioned business calendars and observance policies.
- SQL and native pgx PostgreSQL date codecs with distinct infinity support.
- Clock, temporal, config, validation, wire, and test adapters.
- Exhaustive Gregorian, 19-case mutation, fuzz, race, integration, and
  benchmark gates.
- Blocking allocation budgets for core, business, timezone, wire, and pgx hot
  paths.
- Historical second-offset and date-line timezone vectors plus broad standard
  library differential coverage.
- Shared codec and generated-corpus concurrency proofs, plus an explicit
  business compatibility report.
- Hostile-input fuzzing for every typed calendar parser.
- All-year quarter, semester, and policy-permitted arithmetic inverse proofs.
- Refreshed PostgreSQL 14-18 image pins and actionable integration startup
  diagnostics.
- Portable fuzz and provenance gates without undeclared runner tools.
