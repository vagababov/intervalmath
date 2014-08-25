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

func New(s, e float64) (*Interval, error) {
	if s > e {
		return nil, fmt.Errorf("start(%g) > end(%g)", s, e)
	}
	return &Interval{s, e}, nil
}

func (i *Interval) Positive() bool {
	return i.s > 0
}

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

func (i *Interval) ContainsZero() bool {
	return i.s <= 0 && i.e >= 0
}

func (i *Interval) AddTo(o *Interval) {
	i.s += o.s
	i.e += o.e
}

func Add(a, b *Interval) *Interval {
	return &Interval{a.s + b.s, a.e + b.e}
}
