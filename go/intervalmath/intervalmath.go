// Package intervalmath contains types and functions to operate on interval
// arithmetic.
// See: https://en.wikipedia.org/wiki/Interval_arithmetic
package intervalmath

import (
	"fmt"
	"math"
)

type Interval struct {
	// s is the start of the interval.
	s float64
	// e is the end of the interval.
	e float64
}

// New returns a new Interval.
// error is returned if e < s.
func New(s, e float64) (*Interval, error) {
	if s > e {
		return nil, fmt.Errorf("start(%g) > end(%g)", s, e)
	}
	return &Interval{s, e}, nil
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%g, %g]", i.s, i.e)
}

// Positive returns if interval is completely positive.
func (i *Interval) Positive() bool {
	return i.s > 0
}

// Negative returns if interval is completely negative.
func (i *Interval) Negative() bool {
	return i.e < 0
}

// Equals returns if o is == equal to i. Beware of FP imprecision.
// ApproximatelyEquals compares with given precion.
func Equal(i, o *Interval) bool {
	return i.Equals(o)
}

// ApproximatelyEquals returs if o is approximately equal to i.
func AproximatelyEquals(i, o *Interval, precion float64) bool {
	return i.ApproximatelyEquals(o, precion)
}

// Equals returns if o is == equal to i. Beware of FP imprecision.
// ApproximatelyEquals compares with given precion.
func (i *Interval) Equals(o *Interval) bool {
	return i.s == o.s && i.e == o.e
}

// ApproximatelyEquals returs if o is approximately equal to this.
func (i *Interval) ApproximatelyEquals(o *Interval, precion float64) bool {
	return math.Abs(i.e-o.e) <= precion && math.Abs(i.s-o.s) <= precion
}

// ContainsZero returns true if interval contains 0.
func (i *Interval) ContainsZero() bool {
	return i.s <= 0 && i.e >= 0
}

// Add returns the sum of a and b.
func Add(a, b *Interval) *Interval {
	return &Interval{a.s + b.s, a.e + b.e}
}

// Sub returns difference between a and b.
func Sub(a, b *Interval) *Interval {
	return &Interval{a.s - b.e, a.e - b.s}
}

// Inverse returns the inverse of the interval if it does not contain zero, nil
// will be returned for those intervals.
// For intervals containing zero use two values InverseEx function.
func Inverse(a *Interval) *Interval {
	if a.ContainsZero() {
		return nil
	}
	return &Interval{1. / a.e, 1 / a.s}
}

// InverseEx returns (-inf, 1/a.s] and [1/a.e, +inf). This can be used to
// compute division by interval that contains zero.
func InverseEx(a *Interval) (l, r *Interval) {
	return &Interval{math.Inf(-1), 1. / a.s}, &Interval{1. / a.e, math.Inf(1)}
}
