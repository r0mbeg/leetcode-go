package p0066_plus_1

import (
	"reflect"
	"testing"
)

func Test_plusOne_TableDriven(t *testing.T) {

	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			"test1",
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},

		{
			"test2",
			[]int{4, 3, 2, 1},
			[]int{4, 3, 2, 2},
		},

		{
			"test3",
			[]int{9},
			[]int{1, 0},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			got := plusOne(tt.digits)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})

	}
}
