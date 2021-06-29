// https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_5_A&lang=ja
// OK :)

package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&q)
	M := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&M[i])
	}

	// for _, m := range M {
	res := make([]bool, q)
	for i := 0; i < (1 << uint(n)); i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i & (1 << uint(j))) != 0 {
				sum += a[j]
			}
		}
		for idx := 0; idx < q; idx++ {
			if sum == M[idx] {
				res[idx] = true
			}
		}
	}
	for _, r := range res {
		if r {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
