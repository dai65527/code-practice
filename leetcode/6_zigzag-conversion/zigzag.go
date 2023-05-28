func convert(s string, numRows int) string {
	if s == "" || numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	rowIdx := 0
	direction := -1
	for i := range s {
		rows[rowIdx] += string(s[i])
		if rowIdx == 0 || rowIdx == numRows-1 {
			direction *= -1
		}
		rowIdx += direction
	}
	return strings.Join(rows, "")
}
