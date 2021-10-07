// https://atcoder.jp/contests/pakencamp-2019-day3/tasks/pakencamp_2019_day3_c
// AC :)

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxPt(s1, s2 int, a [][]int, n int) int {
	pt := 0
	for i := 0; i < n; i++ {
		pt += max(a[i][s1], a[i][s2])
	}
	return pt
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	maxpt := 0
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			maxpt = max(maxpt, getMaxPt(i, j, a, n))
		}
	}
	fmt.Println(maxpt)
}
