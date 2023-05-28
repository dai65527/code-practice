type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return 0
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func isValid(s string) bool {
	expectStack := Stack{}
	for _, r := range s {
		if r == '(' {
			expectStack.Push(')')
		} else if r == '[' {
			expectStack.Push(']')
		} else if r == '{' {
			expectStack.Push('}')
		} else {
			if r != expectStack.Pop() {
				return false
			}
		}
	}
	return expectStack.IsEmpty()
}
