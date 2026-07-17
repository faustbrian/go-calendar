# Config, validation, and wire adapters

`calendarconfig.Date` accepts only a string and implements both go-config's
`UnmarshalConfigValue(any)` seam and standard text unmarshalling. Null and
numeric coercion are rejected.

`calendarvalidation.ValidDate` and `DateRange` return dependency-neutral rules.
Wrap a rule in the current `go-validation.ValidatorFunc` and translate its safe
sentinel error into the application's violation code. This avoids a circular or
unpublished module dependency.

`calendarwire.EncodeDate` and `DecodeDate` provide the bounded v1 JSON contract.
The root `Date` also composes directly with go-wire formats that honor standard
text or JSON encoding interfaces.
