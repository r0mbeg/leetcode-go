package p0283_move_zeroes

func moveZeroes(nums []int) {

	i := 0 // non zero ptr

	for j := range nums {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}

	}

}
