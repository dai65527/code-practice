// https://atcoder.jp/contests/abc106/tasks/abc106_b
// AC :)

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 1; i <= n; i += 2 {
		if check(i) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func check(n int) bool {
	cnt := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				cnt++
			} else {
				cnt += 2
			}
		}
	}
	if cnt == 8 {
		return true
	}
	return false
}
