package p0029_divide_two_ints

import "testing"

func Test_divide_TableDriven(t *testing.T) {

	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{
			name:     "trivial_1",
			dividend: 1,
			divisor:  1,
			want:     1,
		},

		{
			name:     "trivial_2",
			dividend: 2,
			divisor:  2,
			want:     1,
		},

		{
			name:     "10 div 3 = 3",
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			name:     "7 div -3 = 2",
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},

		{
			name:     "intMin / -1 = intMax",
			dividend: -2147483648,
			divisor:  -1,
			want:     2147483647,
		},

		{
			name:     "intMax / 2",
			dividend: 2147483647,
			divisor:  2,
			want:     1073741823,
		},
	}

	for _, tt := range tests {

		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res := divide(tt.dividend, tt.divisor)

			if res != tt.want {
				t.Errorf("divide(%d, %d) = %v, want %v", tt.dividend, tt.divisor, res, tt.want)
			}
		})
	}
}
