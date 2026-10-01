# Security threat model and risk register

Version: `CALENDAR-TM-2.0`

Candidate scope updated: 2026-10-01; independent release review pending.

Owner: go-calendar maintainers

## Scope and objectives

This model covers the v2.0.0 root-module source: civil-date and typed-period
construction, parsing, encoding, arithmetic, timezone resolution, business
calendars, PostgreSQL codecs, and the clock, configuration, temporal,
validation, and wire adapters. The deprecated top-level adapter packages are
supported facades over the same boundaries.

The package protects date integrity and process availability by rejecting
non-canonical values, constraining supported years and retained data, requiring
explicit arithmetic and timezone policies, and bounding iterative work. It
does not authenticate holiday providers, authorize business rules, fetch
datasets, update timezone data, open databases, create transactions, emit
telemetry, or start goroutines.

## Assets and trust boundaries

- Civil dates, typed periods, business-day decisions, timezone conversions,
  and PostgreSQL finite/infinity distinctions are integrity-sensitive data.
- Process memory and CPU are availability assets at parser, codec, timezone,
  holiday, metadata, and iterative-search boundaries.
- Holiday names, metadata, provenance, timezone names, and revisions may carry
  application or operational data and must not become secret-bearing or
  attacker-controlled telemetry dimensions.

| Boundary | Trust assumption | Owned control |
| --- | --- | --- |
| Canonical date and typed-period input | caller-controlled | fixed ASCII grammars, exact byte lengths, supported years 1-9999, impossible-value rejection, and bounded diagnostic exposure |
| JSON, text, config, and wire input | caller-controlled | generic JSON admission capped at 64 bytes before decoding, exact ten-byte text admission before conversion, strict date parsing, a separate exact 12-byte canonical wire grammar with a 64-byte cap, and invalid zero-value rejection |
| Arithmetic | caller-controlled integer and policy values | checked multiplication and negation, pre-addition range checks, ordinal day math, sealed policies, and supported-year rejection |
| Timezone name and local time | caller-controlled name and components; trusted runtime tzdata | 255-byte UTF-8 name cap, traversal-like segment rejection, explicit location and gap/fold policy, and at most 145 transition probes |
| Business configuration | application-owned and potentially data-derived | at most seven supplied weekend entries before allocation, bounded holiday/metadata counts and field sizes, length admission before UTF-8 scans, closed weekday/policy enums, deep copies, revision identity, and caller-supplied search limits |
| SQL and pgx values | database-driver input | finite/null/infinity separation, exact date parsing, pre-conversion `[]byte` length rejection, typed codec validation, and no query construction |
| Clock | trusted narrow in-process dependency | explicit capability and location; the package performs one synchronous `Now` call and starts no goroutine |
| Returned values and observability | caller-owned | immutable scalar values, defensive holiday and metadata copies, nominal input-free default diagnostics, and no logging or metrics; explicit decoder-cause inspection retains standard error classification and can expose input-derived details |

## Material risks and dispositions

Every accepted risk has an owner, rationale, mitigation, and review condition.
Severity describes impact under the documented trust assumptions.

| ID | Threat | Severity | Status | Owner, rationale, mitigation, and review condition |
| --- | --- | --- | --- | --- |
| CAL-SEC-001 | Malformed or extreme date and period input causes parser differentials, integer overflow, or invalid civil state. | High | Mitigated | Maintainers own exact ASCII grammars, supported-year validation, checked arithmetic, ordinal bounds, sealed policies, fuzzing, exhaustive Gregorian checks, and mutation tests. Review any parser, representation, range, or arithmetic change. |
| CAL-SEC-002 | Oversized SQL `[]byte` input is copied before rejection and amplifies memory use. | Medium | Fixed | Maintainers reject byte slices longer than the longest accepted SQL token before conversion while preserving accepted values and error identity. The allocation regression exercises a 1 MiB rejected value. Review SQL scanner or token changes. |
| CAL-SEC-003 | Generic date JSON work grows with its representation envelope. | Medium | Mitigated in v2 source | Maintainers own a 64-byte pre-decoder admission cap, preserve legitimate escaped dates and whitespace inside that envelope, and leave receivers unchanged on rejection. The separate wire adapter retains its exact 12-byte grammar. Integrators still bound outer request buffers before they reach this API. Review JSON grammar, envelope limits, decoder changes, or an incident; do not claim the v2 source fixes published v1. |
| CAL-SEC-004 | Timezone gaps, folds, aliases, unusual offsets, or tzdata drift produce an unintended instant. | High | Mitigated subject to trusted tzdata | Maintainers require an explicit location and resolution policy, enumerate verified occurrences with bounded work, and test historical and date-line transitions against the standard library. Deployment owners pin or record tzdata when results must be reproducible. Review timezone logic, supported platforms, tzdata changes, or a conversion incident. |
| CAL-SEC-005 | A crafted zone name causes path traversal or unbounded timezone lookup work. | High | Mitigated | Maintainers cap names before lookup, reject invalid UTF-8, absolute names, backslashes, and empty, dot, or parent segments, then delegate only transition calculation to `time.LoadLocation`. Fuzz and hostile-name tests cover this boundary. Review loading or name-policy changes. |
| CAL-SEC-006 | Mutable or oversized business data changes decisions or creates unbounded validation/search work. | High | Mitigated in v2 source | Maintainers cap holidays, supplied weekend entries, metadata, names, revisions, and provenance fields; reject oversized strings before UTF-8 scans; validate enums; deep-copy inputs and outputs; and require a positive search limit. Existing immutable calendar, closed-calendar, and bounded-input regressions cover these controls; seven/eight-weekday controls characterize the new count admission. Review configuration, observance, copying, validation order, or iteration changes. |
| CAL-SEC-007 | PostgreSQL NULL or infinity is silently reinterpreted as an ordinary date. | High | Mitigated | Maintainers reject NULL and infinity in `Date`, require the explicit `InfinityDate` sum type, validate native pgx states, and exercise SQL/pgx unit, fuzz, and live integration paths. Integrators own query parameters and transaction lifecycle because this package constructs no SQL. Review codec, pgx, or database schema changes. |
| CAL-SEC-008 | Holiday, provenance, revision, or timezone text exposes secrets or creates unbounded observability cardinality. | Medium | Accepted application-data boundary | Application owners choose these display and identity values; the library neither logs nor emits them and a universal semantic allowlist would reject valid v1 data. Keep credentials and personal data out, emit stable revision/provider identifiers instead of dynamic fields, and cap downstream label sets. Review telemetry adoption, new exported fields, or a disclosure incident. |
| CAL-SEC-009 | A malicious or blocking clock implementation panics or stalls `Today`. | Medium | Accepted trusted-capability boundary | Integrators own the in-process clock implementation; arbitrary Go code cannot be safely preempted and recovering its panic would hide corruption. Use a non-nil, bounded, application-owned clock and contain panic at the application boundary. Review new callbacks, clock implementations, or liveness incidents. |
| CAL-SEC-010 | Caller mutation or concurrent use races package state. | High | Mitigated | Maintainers expose immutable value types, copy business maps and slices in both directions, keep no global mutable registry, and run targeted race tests for calendars, codecs, locations, and fixture corpora. Review any pointer, cache, registry, or mutable collection addition. |
| CAL-SEC-011 | Dependency, workflow, or release compromise introduces vulnerable or unreviewed code. | Medium | Mitigated subject to current gates | Maintainers pin module checksums and immutable workflow revisions and run vulnerability, static, license, secret, API, and clean-consumer gates selected for the release boundary. Review every dependency, action, toolchain, suppression, or scanner finding. |
| CAL-SEC-012 | Explicit decoder-cause introspection discloses input-derived diagnostics. | Medium | Accepted explicit introspection boundary | Maintainers own nominal default error formatting without date/period components or decoder-rejected characters. Standard decoder causes remain accessible through `errors.As`/unwrapping to preserve classification. Application owners choose explicit cause inspection and redact it before logging; that choice is not secure default formatting. Review new error wrappers, cause types, raw formatting, observability integration, or disclosure incidents. |

There are no known Critical or open High findings in this model. Residual
Medium risks remain at explicit caller-owned outer buffers, application-data,
trusted in-process clock, and decoder-cause introspection boundaries. They do
not authorize bypass of a package-owned size, state, arithmetic, timezone,
business-search, or persistence control.

## Compatibility, consumers, and release disposition

The historical v1.1.0 API baseline remains unchanged. The active candidate uses
the official `/v2` module/import suffix and an independent v2 API baseline. Direct
owned consumers are go-opening-hours, go-temporal, and the
go-rule-engine temporal adapter. They use root dates, business calendars,
timezone conversion, and deprecated temporal compatibility paths; none depends
on copying oversized rejected PostgreSQL byte values.

Public tags and releases establish v2.0.0 publication status. V1.1.0 does not contain
the SQL allocation repair or the candidate's bounded generic JSON and default
diagnostic policies. The existing SQL repair changes only resource use for
input that already returned `calendar.ErrInvalidFormat`. Generic JSON envelope
rejection and diagnostic changes are intentional v2 contracts, while valid
canonical dates, bounded escaped dates, arithmetic policies, and SQL sentinels
retain their meaning. V1 and v2 date identities are distinct; Opening-hours
aliases and Temporal callers require deliberate consumer migration, not a
silent dependency replacement. See [v2 migration](v2-migration.md). The root
module contains no local replacement directive.

The module is releasable subject to fresh candidate-revision gates and hosted
CI. Scanner results are execution evidence, not durable source claims; a
failed or unavailable required gate blocks a release-ready verdict.

The generic JSON cap applies at the direct `Date.UnmarshalJSON` boundary.
An enclosing `encoding/json` decoder scans its input first and can strip outer
whitespace before invoking this method. Application owners must cap outer
buffers before their decoder; the library does not claim upstream decoder
preemption or acquisition limits. Review decoder composition and ingestion
changes, in addition to the CAL-SEC-003 triggers above.

Review this model after a security incident, before a major release, and when
parsing, arithmetic, timezone, business, persistence, callback, dependency, or
observability boundaries change. Report suspected vulnerabilities through the
private process in [`SECURITY.md`](../SECURITY.md).
