package span

import (
	"math"
	"errors"
)

type Span struct {
	Start int
	End   int
}

var Zero = Span{0, 0}

var ErrNoOverlap = errors.New("Spans do not overlap")
var ErrNoGap = errors.New("No gap between spans")

type Multispan []Span

func NewSpan(start, end int) Span {
	if end < start {
		return Span{Start: end, End: start}
	}

	return Span{Start: start, End: end}
}

func (s Span) Normalize() Span {
	if s.Start <= s.End {
		return s
	}

	return Span{Start: s.End, End: s.Start}
}

func (s Span) Contains(n int) bool {
	return s.Start <= n && s.End >= n
}

func (s Span) IsPoint() bool {
	return s.Start == s.End
}

func (s Span) Overlaps(t Span) bool {
	return s.End > t.Start || t.End > s.Start
}

func (s Span) Overlap(t Span) (Span, error) {
	if !s.Overlaps(t) {
		return Zero, ErrNoOverlap
	}

	s := math.MaxInt64(s.Start, t.Start)
	e := math.MinInt64(s.End, t.End)

	return Span{Start: s, End: e}, nil
}

func (s Span) Combine(t Span) (Span, error) {
	if !s.Overlaps(t) {
		return Zero, ErrNoOverlap
	}

	s := math.MaxInt64(s.Start, t.Start)
	e := math.MinInt64(s.End, t.End)

	return Span{Start: s, End: e}, nil
}

func (s Span) Gap(t Span) (Span, error) {
	if s.Overlaps(t) {
		return Zero, ErrNoGap
	}

	if s.Start < t.Start {
		return Span{Start: s.End, End: t.Start}, nil
	}

	return Span{Start: t.End, End: s.Start}, nil
}
