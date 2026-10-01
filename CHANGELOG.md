# Changelog

All notable changes follow Keep a Changelog. The project uses semantic
versioning.

## Unreleased

## 2.0.0 - Prepared, not published

### Changed

- Use the `github.com/faustbrian/go-calendar/v2` module and import paths on
  main with Go 1.27.0. All canonical adapters and supported compatibility
  packages remain in the same root module; update imports together because
  v1 and v2 civil-date types are distinct.
- Bound generic date JSON at 64 bytes before decoding while retaining
  legitimate whitespace and escaped date strings within that envelope.
  Reject invalid text lengths before conversion and leave receivers unchanged
  on rejection.
- Omit caller-controlled date and period components and decoder-rejected
  characters from default error strings. Error classification remains
  available through `errors.Is` and `errors.As`; explicit decoder-cause
  inspection may reveal input-derived diagnostics.
- Omit unsupported caller type descriptions from configuration and PostgreSQL
  errors without changing rejection, NULL, infinity, or receiver ownership.
- Bound supplied weekend entries at seven before calendar allocation;
  duplicates remain valid within the bound. Reject oversized strings before
  UTF-8 or timezone-name scans, preserving existing classifications.

### Fixed

- Reject oversized PostgreSQL `[]byte` date values before string conversion,
  bounding allocation for invalid driver input without changing accepted date
  and infinity values or their error identities.

### Changed

- Require Go 1.27.0 and record its stable five-allocation wire encode budget;
  the four-allocation decode budget remains unchanged.

## 1.1.0 - 2026-09-09

### Changed

- Add canonical `adapters/clock`, `adapters/config`, `adapters/postgres`,
  `adapters/temporal`, `adapters/validation`, and `adapters/wire` entry points.
  The previously released package paths remain supported deprecated
  compatibility paths with unchanged public contracts.

- Adopt the `go-library-tools` v1.4.0 schema-v2 cohesion contract and local
  `make cohesion` gate without changing calendar API or runtime behavior.
- Pin reusable CI to the immutable v1.4.0 W14-enforcement workflow and enforce
  cohesion metadata in the repository's required CI contract.

- Adopt the versioned shared `golib` repository contract for local and hosted
  verification while retaining package-owned API, provenance, and mutation
  evidence.

### Documentation

- Publish the module's family, capabilities, ownership, lifecycle, supported
  environments, package selection, and delivery status, and link the README to
  the immutable v1.4.0 ecosystem index and family guidance.

- Remove completed implementation plans from the release tree and retain
  package-owned documentation as the maintained reference.

## 1.0.0 - 2026-08-25

### Changed

- Exclude intentional nested modules from root local-proxy archives so local,
  bootstrap, CI, and public module checksums describe the same source
  boundary.

- Track the pinned documentation-tool lockfile so clean CI checkouts install
  the exact validated cspell dependency.

- Reconcile standalone dependency checksums against deterministic current
  module archives so CI, local verification, and release consumers resolve
  identical content.

- Harden standalone documentation validation with deterministic spelling and
  link checks, package-specific documentation gates, and repository-local
  contributor guidance.

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
