func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	ans := 0
	start := 0
	set := map[byte]struct{}{}
	for i, r := range []byte(s) {
		if _, exists := set[r]; exists {
			for s[start] != r {
				delete(set, s[start])
				start++
			}
			start++
		} else {
			set[r] = struct{}{}
			if ans < i-start+1 {
				ans = i - start + 1
			}
		}
	}
	return ans
}
