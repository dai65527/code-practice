package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	digit := 1
	for n > 1000 {
		n /= 10
		digit *= 10
	}
	fmt.Println(n * digit)
}
