func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] >= 10 {
			digits[i] -= 10
			digits[i-1]++
		} else {
			return digits
		}
	}
	if digits[0] >= 10 {
		digits[0] -= 10
		return append([]int{1}, digits...)
	}
	return digits
}
