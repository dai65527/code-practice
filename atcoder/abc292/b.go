package main

import "fmt"

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)
	points := make([]int, n)
	var c, x int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d", &c, &x)
		if c == 1 {
			points[x-1] += 1
		} else if c == 2 {
			points[x-1] += 2
		} else if c == 3 {
			if points[x-1] >= 2 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
