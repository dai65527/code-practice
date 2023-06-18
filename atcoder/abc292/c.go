package main

import "fmt"

func main() {
	var n int
	var ans int
	fmt.Scanf("%d", &n)
	for i := 1; i < n; i++ {
		x := i     // x = ab
		y := n - i // y = cd

		// a*b = xとなるaを探す
		ansX := 0
		for j := 1; j*j <= x; j++ {
			if x%j == 0 {
				ansX++
				// bの分
				if x != j*j {
					ansX++
				}
			}
		}

		// c*d = yとなるcを探す
		ansY := 0
		for j := 1; j*j <= y; j++ {
			if y%j == 0 {
				ansY++
				// bの分
				if y != j*j {
					ansY++
				}
			}
		}
		ans += ansX * ansY
	}
	fmt.Println(ans)
}
