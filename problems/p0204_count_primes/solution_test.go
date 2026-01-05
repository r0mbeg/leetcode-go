package p0204_count_primes

import (
	"testing"
)

func Test_countPrimes_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			in:   10,
			want: 4,
		},
		{
			in:   0,
			want: 0,
		},
		{
			in:   1,
			want: 0,
		},
		{
			in:   2,
			want: 0,
		},
		{
			in:   636381,
			want: 51825,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got := countPrimes(tt.in)

			if got != tt.want {
				t.Fatalf("countPrimes(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
