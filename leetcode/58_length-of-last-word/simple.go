func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			ans++
		} else if ans == 0 {
			continue
		} else {
			return ans
		}
	}
	return ans
}
