package calendar

import "time"

// WeekPolicy is an immutable explicit first-weekday convention. Its zero value
// is invalid even though time.Sunday is numerically zero.
type WeekPolicy struct {
	start time.Weekday
	valid bool
}

// NewWeekPolicy validates and constructs a week policy.
func NewWeekPolicy(start time.Weekday) (WeekPolicy, error) {
	if start < time.Sunday || start > time.Saturday {
		return WeekPolicy{}, ErrInvalidDate
	}
	return WeekPolicy{start: start, valid: true}, nil
}

// IsValid reports whether p was constructed by NewWeekPolicy.
func (p WeekPolicy) IsValid() bool {
	return p.valid && p.start >= time.Sunday && p.start <= time.Saturday
}

// WeekStart returns the configured first weekday.
func (p WeekPolicy) WeekStart() time.Weekday { return p.start }

// StartOfWeek returns the first date of d's policy-defined week. Invalid
// operands or an unrepresentable boundary return an invalid Date, not a clip.
func (p WeekPolicy) StartOfWeek(d Date) Date {
	if !p.IsValid() || !d.IsValid() {
		return Date{}
	}
	offset := (int(d.Weekday()) - int(p.start) + 7) % 7
	start, _ := d.AddDays(-offset)
	return start
}

// EndOfWeek returns the final date of d's policy-defined week, or an invalid
// Date if the complete boundary is outside the supported civil range.
func (p WeekPolicy) EndOfWeek(d Date) Date {
	end, _ := p.StartOfWeek(d).AddDays(6)
	return end
}
