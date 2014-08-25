package intervalmath

import (
	"math"
	"testing"
)

const defaultEpsilon = 0.0001

func TestApeq(t *testing.T) {
  tests := []struct{
    l, r float64
    want bool
  }{
    {0, 0, true},
    {1/1., 2/2., true},
    {1, 2, false},
    {4, 4.000000001, true},
    {3.9999999999, 4, true},
    {math.Inf(-1), math.Inf(-1), true},
    {math.Inf(1), math.Inf(1), true},
    {math.Inf(-1), math.Inf(-1), true},
    {math.Inf(1), 5, false},
    {math.Inf(1)-math.Inf(1), 1, false},
  }
  for _, test := range tests {
    if got, want := apeq(test.r, test.l, defaultEpsilon), test.want; got != want {
      t.Errorf("%g ~ %g: got: %v want: %v", test.r, test.l, got, want)
    }
    if got, want := apeq(test.l, test.r, defaultEpsilon), test.want; got != want {
      t.Errorf("%g ~ %g: got: %v want: %v", test.l, test.r, got, want)
    }
  }
}

func TestCreationAndProps(t *testing.T) {
	tests := []struct {
		desc          string
		s, e          float64
		wantErr       bool
		wantContains0 bool
		wantPositive  bool
		wantNegative  bool
		wantInv           *Interval
	}{
		{"simple", 1.0, 2.0, false, false, true, false, &Interval{0.5, 1.0}},
		{"with zero", -1.0, 1.0, false, true, false, false, nil},
		{"negative", -2, -1, false, false, false, true, &Interval{-1., -0.5}},
		{desc: "bad", s: -1, e: -2, wantErr: true},
	}

	for _, test := range tests {
		i, err := New(test.s, test.e)
		if got, want := err != nil, test.wantErr; got != want {
			t.Errorf("New: err != nil: got: %v want: %v err: %v", got, want, err)
			continue
		}
		// Happy error case.
		if err != nil {
			continue
		}
		if got, want := i.ContainsZero(), test.wantContains0; got != want {
			t.Errorf("%v.ContainsZero: got: %v want: %v", i, got, want)
		}
		if got, want := i.Positive(), test.wantPositive; got != want {
			t.Errorf("%v.Positive: got: %v want: %v", i, got, want)
		}
		if got, want := i.Negative(), test.wantNegative; got != want {
			t.Errorf("%v.Negative: got: %v want: %v", i, got, want)
		}

		// Inverse tests are more complex: can return nil or value.
		if test.wantInv == nil {
      if got, want := Inverse(i), test.wantInv; got != want {
				t.Errorf("1/%v: got: %v want: %v", i, got, want)
			}
			// Test manually for both cases.
			got1, got2 := InverseEx(i)
			want1, want2 := &Interval{math.Inf(-1), 1. / test.s}, &Interval{1. / test.e, math.Inf(1)}
			if !ApproximatelyEqual(got1, want1, defaultEpsilon) {
				t.Errorf("Left 1/%v: got: %v want: %v", i, got1, want1)
			}

			if !ApproximatelyEqual(got2, want2, defaultEpsilon) {
				t.Errorf("Right 1/%v: got: %v want: %v", i, got2, want2)
			}
		} else {
      // without zero, just check equality.
      if got, want := Inverse(i), test.wantInv; !ApproximatelyEqual(got, want, defaultEpsilon) {
        t.Errorf("1/%v: got: %v want: %v", i, got, want)
      }
		}
	}
}
