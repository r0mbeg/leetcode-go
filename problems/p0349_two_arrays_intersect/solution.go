package p0349_two_arrays_intersect

func intersection(nums1 []int, nums2 []int) []int {
	var res []int

	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	countMap1 := make(map[int]int)

	for i := range nums1 {
		//fmt.Printf("nums1[%d]=%d added to map \n", i, nums1[i])
		countMap1[nums1[i]]++
	}

	countMap2 := make(map[int]int)

	for i := range nums2 {
		countMap2[nums2[i]]++
	}

	for k1, _ := range countMap1 {
		if _, ok := countMap2[k1]; ok {
			res = append(res, k1)
		}
	}

	return res
}
