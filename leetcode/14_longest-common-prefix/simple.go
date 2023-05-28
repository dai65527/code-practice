func longestCommonPrefix(strs []string) string {
	preflen := 0
	for len(strs[0]) > preflen {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= preflen || strs[0][preflen] != strs[i][preflen] {
				return strs[0][:preflen]
			}
		}
		preflen++
	}
	return strs[0][:preflen]
}
