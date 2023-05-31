// https://leetcode.com/problems/reverse-words-in-a-string/submissions/960602984/
func reverseWords(s string) string {
	b := []byte(s)
	// remove spaces
	i := 0 // index to write
	j := 0 // start of word
	k := 0 // end of word
	for {
		for j < len(b) && b[j] == ' ' {
			j++
		}
		if j >= len(b) {
			break
		}
		k = j
		for k < len(b) && b[k] != ' ' {
			k++
		}
		if i != 0 {
			b[i] = ' '
			i++
		}
		for j < k {
			b[i] = b[j]
			i++
			j++
		}
	}
	b = b[:i]

	// reverse each words
	i = 0 // start of current word
	j = 0 // end of current word
	for {
		// find this word's end
		for j < len(b) && b[j] != ' ' {
			j++
		}
		// reverse words
		wordlen := j - i
		for k = 0; k < wordlen/2; k++ {
			b[i+k], b[j-k-1] = b[j-k-1], b[i+k]
		}
		// skip space
		if j >= len(b) {
			break
		}
		j++
		i = j
	}

	// reverse entire string
	i = 0
	j = len(b) - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
