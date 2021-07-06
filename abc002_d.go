// https://atcoder.jp/contests/abc002/tasks/abc002_4
// AC :)

package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	x := make([]int, m)
	y := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	maxcount := 1
	for i := 0; i < (1 << n); i++ {
		member := []int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				member = append(member, j+1)
			}
		}
		if isHabatsu(member, x, y) {
			if maxcount < len(member) {
				maxcount = len(member)
			}
		}
	}
	fmt.Println(maxcount)
}

func isHabatsu(member, x, y []int) bool {
	for i := 0; i < len(member); i++ {
		for j := i + 1; j < len(member); j++ {
			if !isFriend(member[i], member[j], x, y) {
				return false
			}
		}
	}
	return true
}

func isFriend(a, b int, x, y []int) bool {
	for i := 0; i < len(x); i++ {
		if a == x[i] && b == y[i] {
			return true
		}
	}
	return false
}
