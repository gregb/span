package span

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestAtoi(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given numeric integer strings", t, func() {

		tests := map[string]int{
			"0":          0,
			"0000":       0,
			"1":          1,
			"20":         20,
			"99":         99,
			"12345":      12345,
			"9876554321": 9876554321,
		}

		Convey("The atoi() fn should return the correct integer version of it", func() {

			for s, i := range tests {
				So(i, ShouldEqual, atoi([]byte(s)))
			}
		})
	})
}

func TestGet(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a know multispan", t, func() {
		s1 := Span{1, 4}
		s2 := Span{3, 6}
		s3 := Span{8, 8}
		ms := Multispan([]Span{s1, s2, s3})

		Convey("The Get() fn should return the correct member spans", func() {

			So(ms.Len(), ShouldEqual, 3)
			So(ms.Get(0), ShouldResemble, s1)
			So(ms.Get(1), ShouldResemble, s2)
			So(ms.Get(2), ShouldResemble, s3)
		})
	})
}

func TestInsertInEmpty(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given an empty multispan", t, func() {
		ms := NewMultiSpan(0)

		Convey("When spans are added", func() {
			s1 := Span{1, 4}
			s2 := Span{3, 6}
			s3 := Span{8, 8}

			ms = ms.Insert(s1, s2, s3)

			Convey("The spans should be present in the multispan in the correct order", func() {

				So(ms.Len(), ShouldEqual, 3)
				So(ms.Get(0), ShouldResemble, s1)
				So(ms.Get(1), ShouldResemble, s2)
				So(ms.Get(2), ShouldResemble, s3)

			})
		})
	})
}

func TestInsertInExisting(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given an existing multispan", t, func() {
		s1 := Span{1, 10}
		s2 := Span{20, 30}
		s3 := Span{40, 50}

		ms := Multispan([]Span{s1, s2, s3})

		Convey("When spans are added", func() {
			i1 := Span{5, 15}
			i2 := Span{30, 35}
			i3 := Span{45, 55}

			ms = ms.Insert(i1, i2, i3)

			Convey("The spans should be present in the multispan correctly interspersed with existing ones", func() {

				So(ms.Len(), ShouldEqual, 6)
				So(ms.Get(0), ShouldResemble, s1)
				So(ms.Get(1), ShouldResemble, i1)
				So(ms.Get(2), ShouldResemble, s2)
				So(ms.Get(3), ShouldResemble, i2)
				So(ms.Get(4), ShouldResemble, s3)
				So(ms.Get(5), ShouldResemble, i3)
			})
		})
	})
}

func TestInsertInExistingLopsidedRight(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given an existing multispan", t, func() {
		s1 := Span{1, 2}
		s2 := Span{4, 5}
		s3 := Span{6, 9}

		ms := Multispan([]Span{s1, s2, s3})

		Convey("When spans are added", func() {
			i1 := Span{10, 15}
			i2 := Span{30, 35}
			i3 := Span{45, 55}

			ms = ms.Insert(i1, i2, i3)

			Convey("The spans should be present in the multispan correctly interspersed with existing ones", func() {

				So(ms.Len(), ShouldEqual, 6)
				So(ms.Get(0), ShouldResemble, s1)
				So(ms.Get(1), ShouldResemble, s2)
				So(ms.Get(2), ShouldResemble, s3)
				So(ms.Get(3), ShouldResemble, i1)
				So(ms.Get(4), ShouldResemble, i2)
				So(ms.Get(5), ShouldResemble, i3)
			})
		})
	})
}

func TestInsertInExistingLopsidedLeft(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given an existing multispan", t, func() {
		s1 := Span{10, 20}
		s2 := Span{40, 50}
		s3 := Span{60, 90}

		ms := Multispan([]Span{s1, s2, s3})

		Convey("When spans are added", func() {
			i1 := Span{1, 2}
			i2 := Span{3, 4}
			i3 := Span{6, 9}

			ms = ms.Insert(i1, i2, i3)

			Convey("The spans should be present in the multispan correctly interspersed with existing ones", func() {

				So(ms.Len(), ShouldEqual, 6)
				So(ms.Get(0), ShouldResemble, i1)
				So(ms.Get(1), ShouldResemble, i2)
				So(ms.Get(2), ShouldResemble, i3)
				So(ms.Get(3), ShouldResemble, s1)
				So(ms.Get(4), ShouldResemble, s2)
				So(ms.Get(5), ShouldResemble, s3)
			})
		})
	})
}

func TestInsertInExistingOutOfOrder(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given an existing multispan", t, func() {
		s1 := Span{1, 10}
		s2 := Span{20, 30}
		s3 := Span{40, 50}

		ms := Multispan([]Span{s1, s2, s3})

		Convey("When spans are added in the wrong order", func() {
			i1 := Span{45, 55}
			i2 := Span{5, 15}
			i3 := Span{30, 35}

			ms = ms.Insert(i1, i2, i3)

			Convey("The spans should be present in the multispan correctly interspersed with existing ones", func() {

				So(ms.Len(), ShouldEqual, 6)
				So(ms.Get(0), ShouldResemble, s1)
				So(ms.Get(1), ShouldResemble, i2)
				So(ms.Get(2), ShouldResemble, s2)
				So(ms.Get(3), ShouldResemble, i3)
				So(ms.Get(4), ShouldResemble, s3)
				So(ms.Get(5), ShouldResemble, i1)
			})
		})
	})
}

func TestNormalizeSimple(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a multispan with overlapping spans", t, func() {

		s1 := Span{1, 4}
		s2 := Span{3, 6}

		ms := Multispan([]Span{s1, s2})

		Convey("When the multispan is normalized", func() {

			ms = ms.Normalize()

			Convey("The spans should be merged", func() {

				want := Span{1, 6}
				So(ms.Len(), ShouldEqual, 1)
				So(ms.Get(0), ShouldResemble, want)

			})
		})
	})
}

func TestNormalizeGapsOnly(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a multispan with overlapping spans", t, func() {

		s1 := Span{1, 4}
		s2 := Span{6, 7}
		s3 := Span{8, 8}

		ms := Multispan([]Span{s1, s2, s3})

		Convey("When the multispan is normalized", func() {

			ms = ms.Normalize()

			Convey("The spans should be merged", func() {

				So(ms.Len(), ShouldEqual, 3)
				So(ms.Get(0), ShouldResemble, s1)
				So(ms.Get(1), ShouldResemble, s2)
				So(ms.Get(2), ShouldResemble, s3)

			})
		})
	})
}

func TestNormalizeMixed(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a multispan with overlapping spans", t, func() {

		s1 := Span{1, 4}
		s2 := Span{2, 7}
		s3 := Span{8, 8}
		s4 := Span{9, 16}
		s5 := Span{11, 13}
		s6 := Span{12, 14}

		ms := Multispan([]Span{s1, s2, s3, s4, s5, s6})

		Convey("When the multispan is normalized", func() {

			ms = ms.Normalize()

			Convey("The spans should be merged", func() {
				w1 := Span{1, 7}
				w2 := Span{8, 8}
				w3 := Span{9, 16}

				So(ms.Len(), ShouldEqual, 3)
				So(ms.Get(0), ShouldResemble, w1)
				So(ms.Get(1), ShouldResemble, w2)
				So(ms.Get(2), ShouldResemble, w3)

			})
		})
	})
}

func TestParse(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given string representations of multispans", t, func() {

		s1 := "1"
		s2 := "1,2"
		s3 := "1-2"
		s4 := "1,2-3"
		s5 := "1-2,3"
		s6 := "12345-67890,1-2,3,4-8"

		Convey("When the strings are parsed", func() {

			ms1, e1 := Parse(s1)
			ms2, e2 := Parse(s2)
			ms3, e3 := Parse(s3)
			ms4, e4 := Parse(s4)
			ms5, e5 := Parse(s5)
			ms6, e6 := Parse(s6)

			Convey("The multispans should be correct", func() {

				a11 := Span{1, 1}
				a12 := Span{1, 2}
				a22 := Span{2, 2}
				a23 := Span{2, 3}
				a33 := Span{3, 3}
				a48 := Span{4, 8}
				abig := Span{12345, 67890}

				w1 := Multispan([]Span{a11})
				w2 := Multispan([]Span{a11, a22})
				w3 := Multispan([]Span{a12})
				w4 := Multispan([]Span{a11, a23})
				w5 := Multispan([]Span{a12, a33})
				w6 := Multispan([]Span{a12, a33, a48, abig})

				So(e1, ShouldBeNil)
				So(ms1, ShouldResemble, w1)

				So(e2, ShouldBeNil)
				So(ms2, ShouldResemble, w2)

				So(e3, ShouldBeNil)
				So(ms3, ShouldResemble, w3)

				So(e4, ShouldBeNil)
				So(ms4, ShouldResemble, w4)

				So(e5, ShouldBeNil)
				So(ms5, ShouldResemble, w5)

				So(e6, ShouldBeNil)
				So(ms6, ShouldResemble, w6)
			})
		})
	})
}
