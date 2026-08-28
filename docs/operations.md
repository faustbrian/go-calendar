# Operations

Pin Go 1.26.6 and the `golib` release declared by `.golib.yaml`. Run
`golib check --all` before release and after OS, container, Go, or tzdata
updates. Review transition
corpus drift rather than weakening assertions.

Monitor classified error counts (`invalid_date`, `nonexistent`, `ambiguous`,
`search_limit`) using bounded labels. Record calendar revision and dataset
checksum in decision logs. Do not record holiday names or hostile input as
telemetry labels.

For database upgrades, run the tagged PostgreSQL integration suite against the target PostgreSQL
version. For business-policy changes, version the calendar, compare decisions,
and retain the prior revision for replay.
