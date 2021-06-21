// https://atcoder.jp/contests/abc095/tasks/arc096_a
// AC :)

package main

import (
	"fmt"
	"math"
)

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var a, b, ab, x, y uint64
	fmt.Scan(&a, &b, &ab, &x, &y)

	price := uint64(math.MaxUint64)
	price = min(price, a*x+b*y)
	if x > y {
		price = min(price, a*(x-y)+ab*2*y)
		price = min(price, ab*2*x)
	} else {
		price = min(price, b*(y-x)+ab*2*x)
		price = min(price, ab*2*y)
	}
	fmt.Println(price)
}
