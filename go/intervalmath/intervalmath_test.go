package intervalmath

import "testing"

func TestCreationAndProps(t *testing.T) {
	tests := []struct {
		desc          string
		s, e          float64
		wantErr       bool
		wantContains0 bool
	}{
		{"simple", 1.0, 2.0, false, false},
		{"with zero", -1.0, 1.0, false, true},
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
			t.Errorf("i.ContainsZero: got: %v want: %v", got, want)
		}
	}
}