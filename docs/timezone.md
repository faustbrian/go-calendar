# Timezones and DST

Civil dates carry no location. An instant is produced only from a
`timezone.LocalDateTime`, an explicit `*time.Location`, and a resolution policy.

- `Reject` accepts exactly one occurrence and rejects both gaps and folds.
- `Earlier` and `Later` select chronological fold occurrences; gaps still fail.
- `MatchOffset(seconds)` selects only an occurrence with that UTC offset.

Resolution samples 145 hourly offsets across ±72 hours around UTC wall
components, then verifies candidates for each distinct sampled offset by round
trip. This is not exhaustive transition traversal for arbitrary locations; an
offset active only between samples is not discovered. Standard-library tzdata
remains authoritative. Explicit LoadLocation("Local") is accepted and can
depend on
process environment; prefer a named IANA identity for reproducible replay.
Applications may import `time/tzdata` when they need an embedded
snapshot; otherwise operating-system updates apply. Persisted local values can
therefore resolve differently after tzdata changes. Persist the zone identity,
policy, intended local value, and tzdata/application version when replay matters.

The compatibility corpus covers New York gap/fold, `US/Eastern` alias,
Lord Howe's 30-minute fold, Kathmandu's +05:45 offset, Helsinki's 23-hour day,
and Apia's 2011 date-line skip. Run `golib check --module .` after any toolchain or
tzdata update and review intentional drift before deployment.
