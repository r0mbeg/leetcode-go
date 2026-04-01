package p0344_reverse_string

func reverseString(s []byte) {

	n := len(s)

	if n <= 1 {
		return
	}

	i := 0
	j := n - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

}
