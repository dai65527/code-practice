// https://atcoder.jp/contests/abc128/tasks/abc128_c
// AC:)

package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([][]int, m)
	k := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&k[i])
		s[i] = make([]int, k[i])
		for j := 0; j < k[i]; j++ {
			fmt.Scan(&s[i][j])
		}
	}
	p := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&p[i])
	}

	ans := 0
	for i := 0; i < (1 << n); i++ {
		yes := true
		for j := 0; j < m; j++ {
			sum := 0
			for l := 0; l < k[j]; l++ {
				if (i & (1 << (s[j][l] - 1))) != 0 {
					sum++
				}
			}
			if sum%2 != p[j] {
				yes = false
				break
			}
		}
		if yes {
			ans++
		}
	}
	fmt.Println(ans)
}
