# Source documentation audit

[Documentation map](index.md) · [Repository README](../README.md)
[Next: API reference](api.md)

## Acceptance and owned boundary

This Tier A pass owns comments and directly contradicted maintained prose at
main a50cfbc48d99e561b78b828724889bc0e17e3be1. Preserve runtime syntax, literals,
directives, dependencies, migration/codec inputs, timezone/Gregorian fixtures,
historical authority and API baselines except the configured current GoDoc.
Consumers are civil arithmetic, application holiday policy, explicit timezone
conversion, SQL/pgx values, config/validation/clock/temporal/wire adapters and
transition fixtures. Named risks are misleading boundary, ambiguity, copying,
mutation, nil/callback and resource claims, not runtime or conformance repair.

## Reproducible before-write inventory

`git ls-files` and `golib source-docs inventory --json` identify one v2 module,
84 Go files: 19 production and 65 tests in 16 package directories. The initial
structural inventory has 20 missing public field/interface comments. Six files
are constrained by integration, race or !race tags. All retained adapter paths,
tests, eleven fuzz targets, eight benchmarks and ten Examples participate in
the audit; build selection does not exclude them. No vendor/generated Go or
local generator entrypoint is present. Transition vectors are maintained source,
not generated source despite returning freshly allocated corpus data. No holiday
dataset is bundled; its native provenance refusal gate remains unchanged.

Native .golib.yaml uses full go-doc mode at api/calendar-v2.txt across all 16
packages. Native golib api update/check owns regeneration; historical
api/baseline.txt is immutable. The declared docs operation runs Examples.
Service-aware Make/CI, full Gregorian/timezone qualification, race, active fuzz,
performance and release checks are not justified by unchanged runtime prose.
No new standards-edition or provider-conformance claim is made.

## Missing versus investigated contracts

Missing comments concern arithmetic component fields, clock methods, business
provenance/config fields and test vectors. Existing core/API docs are substantial;
private arithmetic ordinals, explicit period boundaries, scan-before-copy and
immutable holiday copying need concise nearby rationale, not syntax narration.

Source tracing establishes 0001–9999 proleptic Gregorian civil values, explicit
month policies and invalid zero values. Week boundaries can leave the supported
range and return an invalid Date; ISOWeek validity does not guarantee its Sunday
is representable. Invalid Weekday currently uses zero time, not the documented
Sunday. Equal-date ComponentsUntil returns before policy validation. Date JSON
preserves receivers on error; decoder causes can expose input through errors.As.

Timezone resolution samples 145 hourly offsets across ±72h around UTC wall
components, then round-trip verifies candidates. It is not exhaustive arbitrary
transition enumeration or a heap/time sandbox. Runtime tzdata and explicit Local
location behavior remain deployment-owned. Clock callbacks are synchronous,
without forced cancellation or typed-nil recovery. Adapters own no DB connections,
queries, interval algebra, config registry or goroutine lifecycle. Mutable Scan/
Unmarshal receivers require external synchronization; pgx time inputs use their
civil components without timezone conversion. Both PostgreSQL types reject NULL.

Business calendars deep-copy data and bound selected fields/counts/search steps,
not total application allocation or authoritative provenance. Zero predicates
return false; counting/zero movement can succeed with zero results. Observe
copies effective source data but does not preserve earlier observance flags or
source links, so applying it repeatedly is not idempotent. New observed names can
exceed the input-name bound; output count is checked after constructing copies.

## Smallest sufficient proof

Use pinned Go 1.27.2 formatting and all 84 AST/significant-token/literal/directive
identity; regenerate/inspect the native full GoDoc snapshot and check stability.
Run the declared Example operation, source/local-doc/spec applicability and
pinned changed/staged secret checks, followed by ordinary complete-diff review.
Small deterministic observations may confirm invalid weekday/ISO range and
observance metadata surprises. Keep exact owned three caches/helpers disposable.
Do not execute generators, live PostgreSQL, active fuzz, race, performance or
exhaustive qualification. Coordinator owns remote main, exact-source CI and
aggregate publication; local source proof is not fleet/release completion.


## Outcome and preserved boundaries

Pinned Go 1.27.2 checked all 84 tracked Go files, including the six tagged
files: normalized AST, significant tokens, literals and directives are unchanged.
Raw tokens are also unchanged except the single optional implicit terminator
introduced by expanding FixedClock's documented field; its AST and signature
remain identical. All 388 comment groups parsed and rendered. The final native
full GoDoc snapshot contains all 16 packages; native api update/check passed
and left it stable. Other tracked inputs, including the historical API baseline,
dependency pins, fixtures, dataset policy and dated security review, remain exact.

The declared Example operation passed its ten Examples, and the native holiday
provenance refusal passed. Final prose-only refresh reused these results and
unchanged package renders; no runtime, dependency or example input changed.
Small fixed civil-value observations confirmed invalid Weekday is Monday,
9999-W52 has no representable Sunday, zero-calendar count/zero movement can
succeed, repeated observance resets source flags, and an observed-name overflow
returns no partial output. These observations establish those source contracts,
not a Gregorian/timezone conformance or performance campaign.

Structural source checking reports no findings; local documentation links pass.
Native specification checking finds no applicable standards module or decisions.
Pinned changed-content secret scanning and ordinary final-diff inspection cover
this owned prose boundary. Exact task-owned caches and helpers were removed
immediately after both bounded native sessions. No generator, live database,
race, active fuzz, security qualification or release consumer was executed.
The coordinator retains remote delivery and required exact-source CI ownership.

## Objective Enforcement

The checksum-pinned public Tools v2.4.0 verifier now supplies
`make source-docs-check` as a prerequisite of the existing local and CI Make
routes. Ordinary shared CI selects strict source-documentation checking and
requires its result. Native prerequisites, operation budgets, security and
release controls are retained. Hosted exact-head evidence is recorded by the
coordination ledger; configuration alone does not establish green delivery.
