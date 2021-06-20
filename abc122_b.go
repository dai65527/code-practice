// https://atcoder.jp/contests/abc122/tasks/abc122_b
// AC :)

package main

import "fmt"

func isACGT(c byte) bool {
	return (c == 'A' || c == 'C' || c == 'G' || c == 'T')
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var s string

	fmt.Scan(&s)

	maxlen := 0
	start := 0
	end := 0
	for start = 0; start < len(s); start = end + 1 {
		for end = start; end < len(s) && isACGT(s[end]); end++ {
			// nop
		}
		maxlen = max(maxlen, end-start)
	}
	fmt.Println(maxlen)
}
