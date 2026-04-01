package p0349_two_arrays_intersect

func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	set := make(map[int]struct{})

	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res = append(res, num)
			delete(set, num)
		}
	}

	return res
}
