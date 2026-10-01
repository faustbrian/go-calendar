# Encoding

`Date` text is exactly ten ASCII bytes in `YYYY-MM-DD`. JSON is exactly a JSON
string containing that text. The zero date, `null`, invalid UTF-8, impossible
dates, non-ASCII digits, trailing input, and years outside 0001–9999 fail.

Generic `Date.UnmarshalJSON` accepts whitespace and JSON escapes within the
64-byte `calendar.MaxJSONBytes` envelope, checked before decoding. Text
admission requires exactly `calendar.MaxParseBytes` (ten) bytes before string
conversion. Rejected decoding leaves the receiver unchanged. Default error
text contains no input components; explicit `errors.As` decoder causes may
contain input-derived details and must not be logged without application
redaction.

This cap applies to the bytes passed to `Date.UnmarshalJSON`, not to a larger
request body or a decoder's work before invoking the method. For example,
`encoding/json.Unmarshal` scans its input before calling an unmarshaler and
may omit outer whitespace from that callback. Applications must bound that
outer buffer at acquisition; this method cannot preempt an upstream decoder.

`calendarwire.Version == 1` identifies this stable canonical contract and its
decoder caps input at 64 bytes. Locale-aware display belongs in a presentation
adapter; never use localized display text as a wire value.
