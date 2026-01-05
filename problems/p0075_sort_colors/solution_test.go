package p0075_sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "single_0",
			in:   []int{0},
			want: []int{0},
		},
		{
			name: "single_1",
			in:   []int{1},
			want: []int{1},
		},
		{
			name: "single_2",
			in:   []int{2},
			want: []int{2},
		},
		{
			name: "already_sorted",
			in:   []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "reverse_sorted",
			in:   []int{2, 2, 1, 1, 0, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "mixed_small",
			in:   []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "mixed_with_repeats",
			in:   []int{1, 2, 0, 1, 2, 0, 0, 2, 1},
			want: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			name: "only_zeros",
			in:   []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "only_ones",
			in:   []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "only_twos",
			in:   []int{2, 2},
			want: []int{2, 2},
		},
		{
			name: "two_elements_01",
			in:   []int{0, 1},
			want: []int{0, 1},
		},
		{
			name: "two_elements_10",
			in:   []int{1, 0},
			want: []int{0, 1},
		},
		{
			name: "two_elements_20",
			in:   []int{2, 0},
			want: []int{0, 2},
		},
		{
			name: "two_elements_21",
			in:   []int{2, 1},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// делаем копию входа, чтобы не менять исходные данные тест-кейса
			got := append([]int(nil), tt.in...)

			sortColors(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("sortColors(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func Test_sortColors_InPlaceMutation(t *testing.T) {
	nums := []int{2, 0, 1, 2, 0}
	// Сохраняем указатель на первый элемент (если длина > 0),
	// чтобы убедиться, что работаем с тем же backing array.
	ptrBefore := &nums[0]

	sortColors(nums)

	ptrAfter := &nums[0]
	if ptrBefore != ptrAfter {
		t.Fatalf("expected in-place modification (same backing array), but slice appears to have changed backing storage")
	}

	want := []int{0, 0, 1, 2, 2}
	if !reflect.DeepEqual(nums, want) {
		t.Fatalf("got %v, want %v", nums, want)
	}
}
