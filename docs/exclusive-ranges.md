# Exclusive day ranges

[Documentation map](index.md) · [Repository README](../README.md)
[Next: API reference](api.md)

Never invent `23:59:59.999999999` as an end-of-day value. Civil days vary in
elapsed length and local wall times can be skipped or repeated.

`timezone.DayRange` resolves midnight at the requested date and at the next
date, returning `[start,end)`. `calendartemporal.InclusiveDates` converts
inclusive civil endpoints to the same exclusive instant form suitable for
`temporal/instant.Range`.

Both endpoint midnights must resolve under the chosen policy; there is no
implicit midnight-gap normalization. A range ending at the maximum civil date
cannot manufacture the following exclusive date and returns an error.
