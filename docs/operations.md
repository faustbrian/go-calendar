# Operations

[Documentation map](index.md) · [Repository README](../README.md)
[Next: API reference](api.md)

Use the Go 1.27.2 CI producer and the public `golib` release declared by
`.golib.yaml`; the module minimum remains Go 1.27.0. Run
`make check-full` for relevant release qualification and after OS, container,
Go, or tzdata updates. Review transition corpus drift rather than weakening
assertions.

Monitor classified error counts (`invalid_date`, `nonexistent`, `ambiguous`,
`search_limit`) using bounded labels. Record calendar revision and dataset
checksum in decision logs. Do not record holiday names or hostile input as
telemetry labels.

For database upgrades, run the tagged PostgreSQL integration suite against the target PostgreSQL
version. For business-policy changes, version the calendar, compare decisions,
and retain the prior revision for replay.
