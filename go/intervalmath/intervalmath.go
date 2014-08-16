package intervalmath

import "fmt"

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
