package p0066_plus_1

func plusOne(digits []int) []int {

	vUme := 0

	var res []int

	for i := len(digits) - 1; i >= 0; i-- {

		sas := digits[i] + vUme

		if i == len(digits)-1 {
			sas = sas + 1
		}

		res = append([]int{sas % 10}, res...)

		vUme = sas / 10

	}

	if vUme != 0 {
		res = append([]int{vUme}, res...)
	}

	return res

}
