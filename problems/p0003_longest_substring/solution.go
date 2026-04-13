package p0003_longest_substring

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	result := 0

	left := 0

	seen := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		for {
			if _, ok := seen[s[i]]; !ok {
				break
			}

			delete(seen, s[left])
			left++
			//fmt.Printf("case %s: seen[%s] is a duplicate, left=%d, i=%d, result=%d\n", s, string(s[i]), left, i, result)
		}
		seen[s[i]] = struct{}{}

		if i-left+1 > result {
			result = i - left + 1
		}
		//fmt.Printf("case %s: seen[%s] not duplicate, left=%d, i=%d, result=%d\n", s, string(s[i]), left, i, result)

	}

	return result
}
