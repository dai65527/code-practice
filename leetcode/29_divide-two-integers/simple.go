// https://leetcode.com/problems/divide-two-integers/submissions/960982102/
func divide(dividend int, divisor int) int {
	// d1/d2
	d1 := int32(dividend)
	d2 := int32(divisor)

	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	sign := 1
	if d1 > 0 {
		d1 = -d1
		sign = -sign
	}
	if d2 > 0 {
		d2 = -d2
		sign = -sign
	}
	ans := 0
	for d1 <= d2 {
		d1 = d1 - d2
		ans++
	}
	if sign == -1 {
		return -ans
	}
	return ans
}

