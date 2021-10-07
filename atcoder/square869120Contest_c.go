// https://atcoder.jp/contests/s8pc-6/tasks/s8pc_6_b
// AC :)

package main

import (
	"fmt"
	"math"
)

type route struct {
	a int64
	b int64
}

func main() {
	var n int
	fmt.Scan(&n)

	r := make([]route, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r[i].a, &r[i].b)
	}

	ans := int64(math.MaxInt64)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := int64(0)
			for k := 0; k < n; k++ {
				sum += abs(r[i].a-r[k].a) + abs(r[k].a-r[k].b) + abs(r[k].b-r[j].b)
			}
			ans = min(ans, sum)
		}
	}
	fmt.Println(ans)
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
