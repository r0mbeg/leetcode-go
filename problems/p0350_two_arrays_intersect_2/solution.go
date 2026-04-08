package p0350_two_arrays_intersect_2

func intersect(nums1 []int, nums2 []int) []int {

	nums1 = quicksort(nums1)
	nums2 = quicksort(nums2)

	res := make([]int, 0)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		}
	}

	return res

}

func quicksort(x []int) []int {

	if len(x) < 2 {
		return x
	}

	pivot := x[0]

	var left, right []int

	for _, v := range x[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	result := append(quicksort(left), pivot)
	result = append(result, quicksort(right)...)

	return result

}
