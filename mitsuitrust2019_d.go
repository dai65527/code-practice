// https://atcoder.jp/contests/sumitrust2019/tasks/sumitb2019_d
// AC :)

package main

import "fmt"

func find(s string, c byte, n int, pos int) int {
	for i := pos; i < n; i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	count := 0
	for i := '0'; i <= '9'; i++ {
		pos1 := find(s, byte(i), n, 0)
		if pos1 != -1 {
			for j := '0'; j <= '9'; j++ {
				pos2 := find(s, byte(j), n, pos1+1)
				if pos2 != -1 {
					for k := '0'; k <= '9'; k++ {
						pos3 := find(s, byte(k), n, pos2+1)
						if pos3 != -1 {
							count++
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

// TLE :(
/*
import "fmt"

func key(s string, i, j, k int) int {
	return int(s[i]-'0')*100 + int(s[j]-'0')*10 + int(s[k]-'0')
}

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	passes := map[int]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				passes[key(s, i, j, k)] = 42
			}
		}
	}
	fmt.Println(len(passes))
}
*/
