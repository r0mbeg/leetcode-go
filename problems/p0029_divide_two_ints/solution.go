package p0029_divide_two_ints

const (
	IntMax int64 = 1<<31 - 1
	IntMin int64 = -1 << 31
)

func divide(dividend int, divisor int) int {

	a := int64(dividend)
	b := int64(divisor)

	if a == IntMin && b == -1 {
		return int(IntMax)
	}

	negative := (a < 0) != (b < 0)

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	var res int64 = 0

	for a >= b {
		temp := b
		multiple := int64(1)

		for a >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}

		a -= temp
		res += multiple
	}

	if negative {
		res = -res
	}

	if res > IntMax {
		return int(IntMax)
	}
	if res < IntMin {
		return int(IntMin)
	}
	return int(res)
}
