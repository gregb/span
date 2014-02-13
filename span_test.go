package span

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestNewNormalize(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given two spans with reversed starts and ends created with New", t, func() {
		unnormal := NewSpan(10, 5)
		normal := NewSpan(5, 10)

		Convey("When compared to a reference span known to be normal", func() {
			want := Span{5, 10}

			Convey("The spans should be equal to each other, and the reference", func() {
				So(unnormal, ShouldResemble, want)
				So(normal, ShouldResemble, want)
			})
		})
	})
}

func TestNormalize(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given spans created with object literal notation", t, func() {
		s1 := Span{8, 4}
		s2 := Span{4, 8}

		Convey("When calling Normalize()", func() {
			s1n := s1.Normalize()
			s2n := s2.Normalize()

			want := NewSpan(4, 8)
			Convey("The spans should be normal", func() {
				So(s1n, ShouldResemble, want)
				So(s2n, ShouldResemble, want)
			})
		})
	})
}

func TestPoint(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given point and non-point spans", t, func() {
		s1 := Span{5, 6}
		s2 := Span{6, 6}

		Convey("When calling IsPoint()", func() {
			s1p := s1.IsPoint()
			s2p := s2.IsPoint()

			Convey("The Point-ness should be correctly detected", func() {
				So(s1p, ShouldBeFalse)
				So(s2p, ShouldBeTrue)
			})
		})
	})
}

func TestContains(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a span", t, func() {
		s1 := Span{-2, 8}

		Convey("When testing points in, around, and on the borders", func() {
			cIn := s1.Contains(0)
			cLb := s1.Contains(-2)
			cRb := s1.Contains(8)
			cL := s1.Contains(-1234)
			cR := s1.Contains(33)

			Convey("The Contains() method should produce the correct result", func() {
				So(cIn, ShouldBeTrue)
				So(cLb, ShouldBeTrue)
				So(cRb, ShouldBeTrue)
				So(cL, ShouldBeFalse)
				So(cR, ShouldBeFalse)
			})
		})
	})
}

func TestOverlaps(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given overlapping and non-overlapping spans", t, func() {
		s1 := Span{0, 6}
		s2 := Span{5, 10}
		s3 := Span{8, 8}
		s4 := Span{10, 14}

		//    000000000011111
		//    012345678901234
		// 1: *******--------
		// 2: -----******----
		// 3: --------*------
		// 4: ----------*****

		Convey("When calling Overlaps()", func() {
			o11 := s1.Overlaps(s1)
			o12 := s1.Overlaps(s2)
			o13 := s1.Overlaps(s3)
			o14 := s1.Overlaps(s4)

			o21 := s2.Overlaps(s1)
			o22 := s2.Overlaps(s2)
			o23 := s2.Overlaps(s3)
			o24 := s2.Overlaps(s4)

			o31 := s3.Overlaps(s1)
			o32 := s3.Overlaps(s2)
			o33 := s3.Overlaps(s3)
			o34 := s3.Overlaps(s4)

			o41 := s4.Overlaps(s1)
			o42 := s4.Overlaps(s2)
			o43 := s4.Overlaps(s3)
			o44 := s4.Overlaps(s4)

			Convey("The overlap should be correctly detected", func() {
				So(o11, ShouldBeTrue)
				So(o12, ShouldBeTrue)
				So(o13, ShouldBeFalse)
				So(o14, ShouldBeFalse)

				So(o21, ShouldBeTrue)
				So(o22, ShouldBeTrue)
				So(o23, ShouldBeTrue)
				So(o24, ShouldBeTrue)

				So(o31, ShouldBeFalse)
				So(o32, ShouldBeTrue)
				So(o33, ShouldBeTrue)
				So(o34, ShouldBeFalse)

				So(o41, ShouldBeFalse)
				So(o42, ShouldBeTrue)
				So(o43, ShouldBeFalse)
				So(o44, ShouldBeTrue)

			})
		})
	})
}

func TestOverlap(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given overlapping and non-overlapping spans", t, func() {
		s1 := Span{0, 6}
		s2 := Span{5, 10}
		s3 := Span{8, 8}
		s4 := Span{10, 14}

		Convey("When calling Overlap()", func() {
			o11, e11 := s1.Overlap(s1)
			o12, e12 := s1.Overlap(s2)
			o13, e13 := s1.Overlap(s3)
			o14, e14 := s1.Overlap(s4)

			o21, e21 := s2.Overlap(s1)
			o22, e22 := s2.Overlap(s2)
			o23, e23 := s2.Overlap(s3)
			o24, e24 := s2.Overlap(s4)

			o31, e31 := s3.Overlap(s1)
			o32, e32 := s3.Overlap(s2)
			o33, e33 := s3.Overlap(s3)
			o34, e34 := s3.Overlap(s4)

			o41, e41 := s4.Overlap(s1)
			o42, e42 := s4.Overlap(s2)
			o43, e43 := s4.Overlap(s3)
			o44, e44 := s4.Overlap(s4)

			Convey("The overlap should be correctly computed", func() {

				// all overlap with themselves exactly
				w11 := s1
				w22 := s2
				w33 := s3
				w44 := s4

				// other valid overlaps
				w12 := Span{5, 6}
				w21 := Span{5, 6}
				w23 := Span{8, 8}
				w32 := Span{8, 8}
				w24 := Span{10, 10}
				w42 := Span{10, 10}

				So(o11, ShouldResemble, w11)
				So(e11, ShouldBeNil)
				So(o12, ShouldResemble, w12)
				So(e12, ShouldBeNil)
				So(o13, ShouldResemble, Zero)
				So(e13, ShouldEqual, ErrNoOverlap)
				So(o14, ShouldResemble, Zero)
				So(e14, ShouldEqual, ErrNoOverlap)

				So(o21, ShouldResemble, w21)
				So(e21, ShouldBeNil)
				So(o22, ShouldResemble, w22)
				So(e22, ShouldBeNil)
				So(o23, ShouldResemble, w23)
				So(e23, ShouldBeNil)
				So(o24, ShouldResemble, w24)
				So(e24, ShouldBeNil)

				So(o31, ShouldResemble, Zero)
				So(e31, ShouldEqual, ErrNoOverlap)
				So(o32, ShouldResemble, w32)
				So(e32, ShouldBeNil)
				So(o33, ShouldResemble, w33)
				So(e33, ShouldBeNil)
				So(o34, ShouldResemble, Zero)
				So(e34, ShouldEqual, ErrNoOverlap)

				So(o41, ShouldResemble, Zero)
				So(e41, ShouldEqual, ErrNoOverlap)
				So(o42, ShouldResemble, w42)
				So(e42, ShouldBeNil)
				So(o43, ShouldResemble, Zero)
				So(e43, ShouldEqual, ErrNoOverlap)
				So(o44, ShouldResemble, w44)
				So(e44, ShouldBeNil)

			})
		})
	})
}

func TestCombine(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given overlapping and non-overlapping spans", t, func() {
		s1 := Span{0, 6}
		s2 := Span{5, 10}
		s3 := Span{8, 8}
		s4 := Span{10, 14}

		Convey("When calling Combine()", func() {
			o11, e11 := s1.Combine(s1)
			o12, e12 := s1.Combine(s2)
			o13, e13 := s1.Combine(s3)
			o14, e14 := s1.Combine(s4)

			o21, e21 := s2.Combine(s1)
			o22, e22 := s2.Combine(s2)
			o23, e23 := s2.Combine(s3)
			o24, e24 := s2.Combine(s4)

			o31, e31 := s3.Combine(s1)
			o32, e32 := s3.Combine(s2)
			o33, e33 := s3.Combine(s3)
			o34, e34 := s3.Combine(s4)

			o41, e41 := s4.Combine(s1)
			o42, e42 := s4.Combine(s2)
			o43, e43 := s4.Combine(s3)
			o44, e44 := s4.Combine(s4)

			Convey("The combination should be correctly computed", func() {

				// all combine with themselves exactly
				w11 := s1
				w22 := s2
				w33 := s3
				w44 := s4

				// other valid combinations
				w12 := Span{0, 10}
				w21 := Span{0, 10}
				w23 := Span{5, 10}
				w32 := Span{5, 10}
				w24 := Span{5, 14}
				w42 := Span{5, 14}

				So(o11, ShouldResemble, w11)
				So(e11, ShouldBeNil)
				So(o12, ShouldResemble, w12)
				So(e12, ShouldBeNil)
				So(o13, ShouldResemble, Zero)
				So(e13, ShouldEqual, ErrNoOverlap)
				So(o14, ShouldResemble, Zero)
				So(e14, ShouldEqual, ErrNoOverlap)

				So(o21, ShouldResemble, w21)
				So(e21, ShouldBeNil)
				So(o22, ShouldResemble, w22)
				So(e22, ShouldBeNil)
				So(o23, ShouldResemble, w23)
				So(e23, ShouldBeNil)
				So(o24, ShouldResemble, w24)
				So(e24, ShouldBeNil)

				So(o31, ShouldResemble, Zero)
				So(e31, ShouldEqual, ErrNoOverlap)
				So(o32, ShouldResemble, w32)
				So(e32, ShouldBeNil)
				So(o33, ShouldResemble, w33)
				So(e33, ShouldBeNil)
				So(o34, ShouldResemble, Zero)
				So(e34, ShouldEqual, ErrNoOverlap)

				So(o41, ShouldResemble, Zero)
				So(e41, ShouldEqual, ErrNoOverlap)
				So(o42, ShouldResemble, w42)
				So(e42, ShouldBeNil)
				So(o43, ShouldResemble, Zero)
				So(e43, ShouldEqual, ErrNoOverlap)
				So(o44, ShouldResemble, w44)
				So(e44, ShouldBeNil)

			})
		})
	})
}

func TestGap(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given overlapping and non-overlapping spans", t, func() {
		s1 := Span{0, 6}
		s2 := Span{5, 10}
		s3 := Span{8, 8}
		s4 := Span{10, 14}

		Convey("When calling Gap()", func() {
			o11, e11 := s1.Gap(s1)
			o12, e12 := s1.Gap(s2)
			o13, e13 := s1.Gap(s3)
			o14, e14 := s1.Gap(s4)

			o21, e21 := s2.Gap(s1)
			o22, e22 := s2.Gap(s2)
			o23, e23 := s2.Gap(s3)
			o24, e24 := s2.Gap(s4)

			o31, e31 := s3.Gap(s1)
			o32, e32 := s3.Gap(s2)
			o33, e33 := s3.Gap(s3)
			o34, e34 := s3.Gap(s4)

			o41, e41 := s4.Gap(s1)
			o42, e42 := s4.Gap(s2)
			o43, e43 := s4.Gap(s3)
			o44, e44 := s4.Gap(s4)

			Convey("The gaps should be correctly computed", func() {

				// valid gaps
				w13 := Span{6, 8}
				w31 := Span{6, 8}
				w14 := Span{6, 10}
				w41 := Span{6, 10}
				w34 := Span{8, 10}
				w43 := Span{8, 10}

				So(o11, ShouldResemble, Zero)
				So(e11, ShouldEqual, ErrNoGap)
				So(o12, ShouldResemble, Zero)
				So(e12, ShouldEqual, ErrNoGap)
				So(o13, ShouldResemble, w13)
				So(e13, ShouldBeNil)
				So(o14, ShouldResemble, w14)
				So(e14, ShouldBeNil)

				So(o21, ShouldResemble, Zero)
				So(e21, ShouldEqual, ErrNoGap)
				So(o22, ShouldResemble, Zero)
				So(e22, ShouldEqual, ErrNoGap)
				So(o23, ShouldResemble, Zero)
				So(e23, ShouldEqual, ErrNoGap)
				So(o24, ShouldResemble, Zero)
				So(e24, ShouldEqual, ErrNoGap)

				So(o31, ShouldResemble, w31)
				So(e31, ShouldBeNil)
				So(o32, ShouldResemble, Zero)
				So(e32, ShouldEqual, ErrNoGap)
				So(o33, ShouldResemble, Zero)
				So(e33, ShouldEqual, ErrNoGap)
				So(o34, ShouldResemble, w34)
				So(e34, ShouldBeNil)

				So(o41, ShouldResemble, w41)
				So(e41, ShouldBeNil)
				So(o42, ShouldResemble, Zero)
				So(e42, ShouldEqual, ErrNoGap)
				So(o43, ShouldResemble, w43)
				So(e43, ShouldBeNil)
				So(o44, ShouldResemble, Zero)
				So(e44, ShouldEqual, ErrNoGap)

			})
		})
	})
}
