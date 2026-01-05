package p0045_jump_game_2

import "testing"

func Test_jump_TableDriven(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			nums: []int{1, 2, 3},
			want: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res := jump(tt.nums)
			if res != tt.want {
				t.Errorf("jump(%v) = %v,want %v", tt.nums, res, tt.want)
			}
		})
	}
}
