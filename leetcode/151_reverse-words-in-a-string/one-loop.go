func reverseWords(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	j := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if i == j-1 {
				j--
			} else {
				if sb.Len() > 0 {
					sb.WriteString(" ")
				}
				sb.WriteString(s[i+1 : j])
				j = i
			}
		}
	}
	if j != 0 {
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(s[:j])
	}
	return sb.String()
}
