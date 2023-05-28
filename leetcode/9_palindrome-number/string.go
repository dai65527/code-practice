func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	n := len(s) / 2
	for i := 0; i < n; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
