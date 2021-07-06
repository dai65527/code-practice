// https://atcoder.jp/contests/joi2008yo/tasks/joi2008yo_e
// AC :)

package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	var in int
	s := make([]int, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&in)
			if in != 0 {
				s[j] |= (1 << i)
			}
		}
	}

	maxcount := 0
	for i := 0; i < (1 << r); i++ {
		count := 0
		for j := 0; j < c; j++ {
			sj := s[j]
			for k := 0; k < r; k++ {
				if (i & (1 << k)) != 0 {
					sj ^= (1 << k)
				}
			}
			count += maxsenbei(sj, r)
		}
		if count > maxcount {
			maxcount = count
		}
	}

	fmt.Println(maxcount)
}

func maxsenbei(s, r int) int {
	n := 0
	m := 0
	for i := 0; i < r; i++ {
		if (s & (1 << i)) != 0 {
			n++
		} else {
			m++
		}
	}
	if n > m {
		return n
	}
	return m
}
