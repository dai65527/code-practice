func reverse(x int) int {
	// "Assume the environment does not allow you to store 64-bit integers (signed or unsigned)."
	v := int32(x)
	if v == math.MinInt32 {
		return 0
	}
	sign := 1
	if v < 0 {
		sign = -1
		v = -v
	}
	ans := int32(0)
	for v > 0 {
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && v%10 > math.MaxInt32%10) {
			return 0
		}
		ans = ans*10 + v%10
		v /= 10
	}
	return sign * int(ans)
}
