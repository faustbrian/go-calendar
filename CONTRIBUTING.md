# Contributing

## Before Editing

1. Read [`AGENTS.md`](AGENTS.md) and the affected package's docs.
2. Run `go test ./...` or the affected package's narrower tests.
3. Identify owned dependencies and reverse dependants in `modules.json`.
4. Preserve unrelated work and generated/corpus provenance.

## Changes

Keep commits focused and conventional. Update every affected changelog with
the behavior and migration impact. Public API changes require compatibility
evidence and documentation. Specification behavior requires a decision record,
fixture coverage, and interoperability evidence.

New direct dependencies and dependency updates must follow the
[dependency governance policy](AGENTS.md#dependencies-and-supply-chain). Package-local
update bots are forbidden; the root policy owns every module and action update.

Specification-backed changes must follow the
[specification governance contract](AGENTS.md#design), update
the affected stable decision entries, and complete the Specification Decisions
section of the pull request template. An unresolved interpretation or stale
source pin is release-blocking; peer behavior cannot silently select policy.

Mutation is optional local or scheduled assurance, not a required PR status.
Confirmed behavioral and security defects still require a focused regression
and repair; equivalent/invalid mutants retain reviewed narrow dispositions.

Do not add package-local workflows, permanent replacements, machine-specific
paths, bypass flags, broad mutation exclusions, or aggregate quality metrics
that hide a failing package.

## Verification

Standalone `go test ./...` runs deterministic tests, fuzz seeds and saved corpus
regressions. It needs no sibling checkout or ecosystem goal ledger. Tagged live
PostgreSQL tests are separate from this untagged Go route.

With the checksum-pinned public verifier selected by `.golib.yaml`, use:

```bash
make inventory
make check
make ci
```

`check` and `ci` use bounded local checks; ordinary CI consumes the public
binary rather than building the verifier from source. The manifest requires
PostgreSQL and the `integration` test tag, so the service-aware Make route may
require Docker and is not equivalent to untagged Go tests. No new ordinary race
campaign is required for the immutable core. Public API, timezone, persistence
and adapter changes retain their affected compatibility, integration, corpus
and direct-consumer checks. Preserve the holiday-dataset provenance contract.

Short active exploration selects one target. Optional local mutation selects
one package; neither command is an all-change prerequisite:

```bash
GOMAXPROCS=2 go test -run '^$' -fuzz '^FuzzParseDate$' -fuzztime=2s -parallel=1 -timeout=2m .
golib mutation --module . --package ./business
```

For PostgreSQL changes, use a disposable database and run both the canonical
and retained adapters. Never use a shared or production database:

```bash
POSTGRES_URL='postgres://postgres:postgres@127.0.0.1:5432/calendar?sslmode=disable' \
go test -tags=integration -run '^TestPostgreSQLDateRoundTrip$' -count=1 -timeout=5m ./adapters/postgres ./postgres
```

`make check-full` preserves explicit full assurance, including the timezone
selection, provenance refusal, all eleven native fuzz budgets and benchmarks.
Relevant release qualification uses `golib release dry-run --module .`; the
manual release workflow additionally retains its OS and PostgreSQL matrices.
The selected security diagnostic remains development qualification, not release
substitution. Report unavailable, skipped and failed boundaries honestly; a
local pass is not full assurance, hosted-CI or public-release qualification.

## Adding A Module

Follow [repository structure policy](AGENTS.md#repository-structure). New modules
require an explicit purpose, ownership boundary, dependency review, package
catalog entry, full quality gates, documentation, changelog, license, security
policy, compatibility plan, and release dry-run.

## Source Documentation

Install the checksum-pinned public verifier declared in `.golib.yaml`.
`make source-docs-check` validates objective Go comment, marker, link and
generated-producer structure; `make check`, `make check-full` and `make ci`
require it before their existing checks. Shared CI runs the same strict
contract and includes its result in Required. This is structural enforcement,
not a substitute for source-truth review or behavioral verification.
