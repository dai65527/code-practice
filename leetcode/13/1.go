type Pair struct {
	Str string
	Num int
}

var pairs []Pair = []Pair{
	{"IV", 4},
	{"IX", 9},
	{"XL", 40},
	{"XC", 90},
	{"CD", 400},
	{"CM", 900},
	{"I", 1},
	{"V", 5},
	{"X", 10},
	{"L", 50},
	{"C", 100},
	{"D", 500},
	{"M", 1000},
}

func romanToInt(s string) int {
	n := len(s)
	ans := 0
	i := 0
	for i < n {
		for _, pair := range pairs {
			if strings.HasPrefix(s[i:], pair.Str) {
				ans += pair.Num
				i += len(pair.Str)
				break
			}
		}
	}
	return ans
}
