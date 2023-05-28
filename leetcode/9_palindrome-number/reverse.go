func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	tmp := x
	for tmp > 0 {
		rev = rev*10 + tmp%10
		tmp = tmp / 10
	}
	return x == rev
}
