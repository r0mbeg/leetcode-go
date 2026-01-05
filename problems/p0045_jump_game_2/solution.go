package p0045_jump_game_2

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 0
	end := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == end {
			jumps++
			end = farthest
		}
	}
	return jumps
}
