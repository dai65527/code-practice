var vmap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	vb := 0
	ans := 0
	for _, c := range s {
		vc := vmap[c]
		if vb != 0 && vb < vc {
			ans += vc - vb
			vb = 0
		} else {
			ans += vb
			vb = vc
		}
	}
	return ans + vb
}
