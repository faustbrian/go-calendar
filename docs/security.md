# Security threat model and risk register

Version: `CALENDAR-TM-1.0`

Reviewed: 2026-09-13

Owner: go-calendar maintainers

## Scope and objectives

This model covers the released v1 root module: civil-date and typed-period
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
| JSON, text, config, and wire input | caller-controlled within its decoder envelope | strict date parsing, a 64-byte canonical wire cap, and invalid zero-value rejection |
| Arithmetic | caller-controlled integer and policy values | checked multiplication and negation, pre-addition range checks, ordinal day math, sealed policies, and supported-year rejection |
| Timezone name and local time | caller-controlled name and components; trusted runtime tzdata | 255-byte UTF-8 name cap, traversal-like segment rejection, explicit location and gap/fold policy, and at most 145 transition probes |
| Business configuration | application-owned and potentially data-derived | bounded counts and field sizes, UTF-8 validation, closed weekday/policy enums, deep copies, revision identity, and caller-supplied search limits |
| SQL and pgx values | database-driver input | finite/null/infinity separation, exact date parsing, pre-conversion `[]byte` length rejection, typed codec validation, and no query construction |
| Clock | trusted narrow in-process dependency | explicit capability and location; the package performs one synchronous `Now` call and starts no goroutine |
| Returned values and observability | caller-owned | immutable scalar values, defensive holiday and metadata copies, typed errors that never retain payload objects, and no logging or metrics; some validation errors include bounded numeric components or one decoder-rejected character |

## Material risks and dispositions

Every accepted risk has an owner, rationale, mitigation, and review condition.
Severity describes impact under the documented trust assumptions.

| ID | Threat | Severity | Status | Owner, rationale, mitigation, and review condition |
| --- | --- | --- | --- | --- |
| CAL-SEC-001 | Malformed or extreme date and period input causes parser differentials, integer overflow, or invalid civil state. | High | Mitigated | Maintainers own exact ASCII grammars, supported-year validation, checked arithmetic, ordinal bounds, sealed policies, fuzzing, exhaustive Gregorian checks, and mutation tests. Review any parser, representation, range, or arithmetic change. |
| CAL-SEC-002 | Oversized SQL `[]byte` input is copied before rejection and amplifies memory use. | Medium | Fixed | Maintainers reject byte slices longer than the longest accepted SQL token before conversion while preserving accepted values and error identity. The allocation regression exercises a 1 MiB rejected value. Review SQL scanner or token changes. |
| CAL-SEC-003 | Generic `Date.UnmarshalJSON` performs work proportional to the caller-provided JSON slice, including whitespace or escaped representations accepted by released v1. | Medium | Accepted decoder-envelope boundary | Integrators own the upstream request or message byte limit because v1 has accepted unbounded JSON whitespace and changing that behavior would break the released decoder contract. Use `adapters/wire.DecodeDate` for hostile canonical wire input; it accepts only the exact 12-byte representation and rejects above 64 bytes before decoding. Review every JSON boundary incident and design a strict cap only on a future unpublished `/v2` line with an independent API baseline and consumer migration evidence. |
| CAL-SEC-004 | Timezone gaps, folds, aliases, unusual offsets, or tzdata drift produce an unintended instant. | High | Mitigated subject to trusted tzdata | Maintainers require an explicit location and resolution policy, enumerate verified occurrences with bounded work, and test historical and date-line transitions against the standard library. Deployment owners pin or record tzdata when results must be reproducible. Review timezone logic, supported platforms, tzdata changes, or a conversion incident. |
| CAL-SEC-005 | A crafted zone name causes path traversal or unbounded timezone lookup work. | High | Mitigated | Maintainers cap names before lookup, reject invalid UTF-8, absolute names, backslashes, and empty, dot, or parent segments, then delegate only transition calculation to `time.LoadLocation`. Fuzz and hostile-name tests cover this boundary. Review loading or name-policy changes. |
| CAL-SEC-006 | Mutable or oversized holiday data changes concurrent decisions, exhausts memory, or creates an infinite business-day search. | High | Mitigated | Maintainers cap holidays, metadata, names, revisions, and provenance fields; validate enums and UTF-8; deep-copy inputs and outputs; require a positive search limit; and provide race, fuzz, closed-calendar, and resource tests. Review business configuration, observance, copying, or iteration changes. |
| CAL-SEC-007 | PostgreSQL NULL or infinity is silently reinterpreted as an ordinary date. | High | Mitigated | Maintainers reject NULL and infinity in `Date`, require the explicit `InfinityDate` sum type, validate native pgx states, and exercise SQL/pgx unit, fuzz, and live integration paths. Integrators own query parameters and transaction lifecycle because this package constructs no SQL. Review codec, pgx, or database schema changes. |
| CAL-SEC-008 | Holiday, provenance, revision, or timezone text exposes secrets or creates unbounded observability cardinality. | Medium | Accepted application-data boundary | Application owners choose these display and identity values; the library neither logs nor emits them and a universal semantic allowlist would reject valid v1 data. Keep credentials and personal data out, emit stable revision/provider identifiers instead of dynamic fields, and cap downstream label sets. Review telemetry adoption, new exported fields, or a disclosure incident. |
| CAL-SEC-009 | A malicious or blocking clock implementation panics or stalls `Today`. | Medium | Accepted trusted-capability boundary | Integrators own the in-process clock implementation; arbitrary Go code cannot be safely preempted and recovering its panic would hide corruption. Use a non-nil, bounded, application-owned clock and contain panic at the application boundary. Review new callbacks, clock implementations, or liveness incidents. |
| CAL-SEC-010 | Caller mutation or concurrent use races package state. | High | Mitigated | Maintainers expose immutable value types, copy business maps and slices in both directions, keep no global mutable registry, and run targeted race tests for calendars, codecs, locations, and fixture corpora. Review any pointer, cache, registry, or mutable collection addition. |
| CAL-SEC-011 | Dependency, workflow, or release compromise introduces vulnerable or unreviewed code. | Medium | Mitigated subject to current gates | Maintainers pin module checksums and immutable workflow revisions and run vulnerability, static, license, secret, API, and clean-consumer gates selected for the release boundary. Review every dependency, action, toolchain, suppression, or scanner finding. |
| CAL-SEC-012 | Validation diagnostics disclose caller-supplied civil-date components. | Medium | Accepted bounded diagnostic boundary | Maintainers preserve released v1 diagnostics because constructor errors include only bounded numeric year, month, day, quarter, semester, or week components, and standard JSON syntax errors may identify one rejected character; the package never logs or retains the source payload. Integrators must treat validation errors as application data and avoid logging them where date components are sensitive. Review any new error field, raw-string formatting, observability integration, disclosure incident, or a future v2 error-contract design. |

There are no known Critical or open High findings in this model. Residual
Medium risks remain at explicit caller-owned JSON envelope, application-data,
trusted in-process clock, and bounded validation-diagnostic boundaries. They do
not authorize bypass of a package-owned size, state, arithmetic, timezone,
business-search, or persistence control.

## Compatibility, consumers, and release disposition

The released v1.1.0 API baseline is unchanged on the active v1 line. Direct
owned consumers are go-opening-hours, go-temporal, and the
go-rule-engine temporal adapter. They use root dates, business calendars,
timezone conversion, and deprecated temporal compatibility paths; none depends
on copying oversized rejected PostgreSQL byte values.

There is no v2 module or v2 release claim. The SQL allocation repair changes
only resource use for input that already returned `calendar.ErrInvalidFormat`.
Any future strict JSON envelope or other breaking security behavior must be
designed on an unpublished `/v2` line with an immutable v1 baseline, an active
v2 API baseline, and consumer migration evidence. The root module contains no
local replacement directive.

The module is releasable subject to fresh candidate-revision gates and hosted
CI. Scanner results are execution evidence, not durable source claims; a
failed or unavailable required gate blocks a release-ready verdict.

Review this model after a security incident, before a major release, and when
parsing, arithmetic, timezone, business, persistence, callback, dependency, or
observability boundaries change. Report suspected vulnerabilities through the
private process in [`SECURITY.md`](../SECURITY.md).
