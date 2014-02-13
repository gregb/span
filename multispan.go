package span

import (
	"sort"
)

type Multispan []Span

func NewMultiSpan(cap int) Multispan {
	return make([]Span, 0, cap)
}

func (ms Multispan) Insert(spans ...Span) Multispan {

	sorted := make([]Span, 0, len(spans)+len(ms))

	sort.Sort(Multispan(spans))

	if len(ms) == 0 {
		return spans
	}

	mi := 0
	si := 0

	for mi < len(ms) {

		if si > len(spans)-1 {
			sorted = append(sorted, ms[mi])
			mi++
			continue
		}

		existing := ms[mi]
		toInsert := spans[si]

		if toInsert.Start < existing.Start {
			sorted = append(sorted, toInsert)
			si++
		} else {
			sorted = append(sorted, existing)
			mi++
		}
	}

	sorted = append(sorted, spans[si:]...)

	return sorted
}

func (ms Multispan) Normalize() Multispan {

	if len(ms) == 1 {
		return ms
	}

	sort.Sort(ms)

	spans := make([]Span, 0, len(ms))

	var left, right, combined Span
	var err error

	left = ms[0]
	for i := 1; i < len(ms); i++ {

		right = ms[i]

		combined, err = left.Combine(right)

		if err == nil {
			left = combined

		} else {
			spans = append(spans, left)
			left = right
		}

		if i+1 == len(ms) {
			spans = append(spans, left)
		}
	}

	return spans
}

func (ms Multispan) Get(i int) Span {
	return ms[i]
}

// implements sort.Interface
// Len is the number of elements in the collection.
func (ms Multispan) Len() int {
	return len(ms)
}

// implements sort.Interface
// Less reports whether the element with
// index i should sort before the element with index j.
func (ms Multispan) Less(i, j int) bool {
	return ms[i].Start < ms[j].Start
}

// implements sort.Interface
// Swap swaps the elements with indexes i and j.
func (ms Multispan) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func Parse(s string) (Multispan, error) {

	ms := Multispan{}

	bytes := []byte(s)

	n := -1
	mark := 0

	for i, b := range bytes {

		if b == ',' {
			t := atoi(bytes[mark:i])

			if n > 0 {
				// there was already a parsed number here since the last comma
				// so this is a range
				ms = ms.Insert(Span{n, t})
			} else {
				// no other numbers since last comma
				// this is a point range
				ms = ms.Insert(Span{t, t})
			}

			// reset n and mark the next spot
			n = -1
			mark = i + 1
		}

		if i == len(bytes)-1 {
			t := atoi(bytes[mark : i+1])

			if n > 0 {
				// there was already a parsed number here since the last comma
				// so this is a range
				ms = ms.Insert(Span{n, t})
				n = -1
			} else {
				// no other numbers since last comma
				// this is a point range
				ms = ms.Insert(Span{t, t})
			}
		}

		if b == '-' {
			n = atoi(bytes[mark:i])
			mark = i + 1
		}
	}

	return ms, nil
}

// Dumb, fast string to int converter
// Restrictions: ascii only, base 10 only, positive numbers only, no overflow check
func atoi(b []byte) int {

	n := 0
	p := 1

	for i := len(b) - 1; i >= 0; i-- {
		v := int(b[i] - 48)
		n += v * p
		p *= 10
	}

	return n
}
