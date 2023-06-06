package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%dj", &n)
	s := make([]string, n)
	a := make([]int, n)

	minIndex := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &s[i], &a[i])
		if a[i] < a[minIndex] {
			minIndex = i
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(s[(minIndex+i)%n])
	}
}
