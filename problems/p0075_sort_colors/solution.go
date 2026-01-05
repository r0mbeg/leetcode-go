package p0075_sort_colors

func sortColors(nums []int) {

	redCount := 0
	whiteCount := 0
	blueCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			redCount++
		} else if nums[i] == 1 {
			whiteCount++
		} else if nums[i] == 2 {
			blueCount++
		}
	}

	for i := 0; i < redCount; i++ {
		nums[i] = 0
	}

	for i := redCount; i < redCount+whiteCount; i++ {
		nums[i] = 1
	}

	for i := redCount + whiteCount; i < redCount+whiteCount+blueCount; i++ {
		nums[i] = 2
	}

}
