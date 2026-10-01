# V2 migration

This guide describes v2.0.0 source on main. Public tags and GitHub releases
establish publication status; do not infer publication from this guide.
Production source remains at the repository root, without version directories
or branches. Published v1.1.0 remains a distinct historical module identity.

When the release is published, require `github.com/faustbrian/go-calendar/v2`
and change every owning Calendar import to that prefix. The 16 root-module
packages, including supported deprecated compatibility paths, remain available.
There are no separately versioned Calendar adapters to migrate independently.
Go 1.27.0 is required.

V1 and v2 `Date`, business-calendar, and timezone values have distinct Go
identities. Upgrade public adapters and direct consumers coherently; an alias
such as Opening-hours `Date` changes its identity when its Calendar dependency
changes. Temporal callers and their reverse consumers need their own
compatibility assessment. This release does not claim those external changes
have already occurred, and does not rewrite historical ecosystem cohorts.

Generic date JSON now has a 64-byte envelope, including whitespace and
escapes. Canonical dates and all legitimate escaped representations inside
that envelope retain their meaning. Previously accepted larger whitespace
envelopes are rejected. The strict wire adapter remains different: it accepts
only the exact 12-byte canonical JSON date. Text continues to require ten
ASCII bytes. Rejections do not replace a valid receiver.

The JSON cap belongs to the byte slice passed directly to `Date.UnmarshalJSON`.
An enclosing `encoding/json` decoder can scan a larger document first and omit
outer whitespace when invoking this method. Bound the whole request or document
before that decoder; the Date method does not own upstream acquisition.

Default constructor, period, arithmetic, and JSON-decoder error strings no
longer reveal input-derived date components or rejected characters. Do not
parse human-readable error strings; use `errors.Is` for nominal Calendar
classification and `errors.As` for explicit decoder classification. Inspecting
or formatting the decoder cause is an application-selected data exposure
boundary: it may still describe rejected input. The library does not erase
caller buffers or redact application logging of causes.

Unsupported config and PostgreSQL values still reject without changing a
receiver, but error text no longer includes the caller's type description.
NULL, infinity, and date-format classifications remain distinct.

Business configuration accepts at most `business.MaxWeekends` (seven) supplied
weekday entries, counted before calendar allocation. Duplicate entries remain
valid within that bound and have their existing set meaning; larger duplicate
lists now return `business.ErrResourceLimit`. Invalid revision retains its
first validation priority and invalid weekdays within the bound still return
`business.ErrInvalidCalendar`. Revision, metadata, and timezone-name byte caps
now short-circuit before UTF-8 and name scans, without changing accepted text.

SQL finite, NULL, and infinity distinctions, civil-date arithmetic, canonical
encoding, business-calendar ownership, and explicit DST policies are retained.
The historical v1 API baseline remains immutable; the active v2 baseline is
`api/calendar-v2.txt`. Publication still requires exact-source CI, release
rehearsal, verified artifacts, and a clean public consumer.
