package p0350_two_arrays_intersect_2

import (
	"fmt"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {

	testcase := []struct {
		nums1        []int
		nums2        []int
		intersection []int
	}{

		{
			nums1:        []int{},
			nums2:        []int{},
			intersection: []int{},
		},

		{
			nums1:        []int{1, 3, 4, 5, 6, 7},
			nums2:        []int{7},
			intersection: []int{7},
		},

		{
			nums1:        []int{1},
			nums2:        []int{},
			intersection: []int{},
		},

		{
			nums1:        []int{1, 2, 2, 1},
			nums2:        []int{2, 2},
			intersection: []int{2, 2},
		},

		{
			nums1:        []int{4, 9, 5},
			nums2:        []int{9, 4, 9, 8, 4},
			intersection: []int{4, 9},
		},
	}

	for i, testcase := range testcase {

		got := intersect(testcase.nums1, testcase.nums2)

		if len(got) != len(testcase.intersection) {
			t.Errorf("case %d incorrect: got %v, want %v", i, got, testcase.intersection)
		}

		sort.Ints(got)
		sort.Ints(testcase.intersection)

		for i := 0; i < len(got); i++ {
			if got[i] != testcase.intersection[i] {
				t.Errorf("case %d incorrect: got %v, want %v", i, got, testcase.intersection)
			}
		}

		fmt.Printf("case %d correct: got %#v\n", i, got)

	}

}
