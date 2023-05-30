func strStr(haystack string, needle string) int {
	for i := range haystack {
		ok := true
		for j := range needle {
			if i+j >= len(haystack) || haystack[i+j] != needle[j] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}
