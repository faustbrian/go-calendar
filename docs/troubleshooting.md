# Troubleshooting

## `ErrInvalidDate`

Check the 1–9999 range, Gregorian month length, and zero values. Parsing accepts
only `YYYY-MM-DD` ASCII text.

## `ErrNonexistent` or `ErrAmbiguous`

The local wall value crosses a timezone transition. Choose a domain policy;
do not retry with a random location or fabricate an offset.

## `ErrSearchLimit`

The calendar is closed longer than the supplied budget or the budget is zero.
Inspect weekend/holiday configuration before increasing it.

## Timezone results changed

Compare Go, OS/container tzdata, zone identity, and stored policy versions. Run
`golib check --module .` and review the corpus.

## Integration cannot start

The shared `golib` service fixture starts an isolated PostgreSQL instance for
`golib check --module .`. Never point verification at a production database.
