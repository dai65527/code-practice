// https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ITP1_7_B&lang=ja
// passed

package main

import "fmt"

func main() {
	var n, x []int
	var n_n, x_n int
	var cnt int

	for cnt = 0; ; cnt++ {
		fmt.Scan(&n_n, &x_n)
		if n_n == 0 && x_n == 0 {
			break
		}
		n = append(n, n_n)
		x = append(x, x_n)
	}

	for i := 0; i < len(n); i++ {
		find(n[i], x[i])
	}
}

func find(n, x int) {
	cnt := 0
	for i := 1; i <= n && i < x; i++ {
		for j := i + 1; j <= n && i+j < x; j++ {
			for k := j + 1; k <= n && i+j+k <= x; k++ {
				if i+j+k == x {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
