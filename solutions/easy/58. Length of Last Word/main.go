func lengthOfLastWord(s string) int {
	lenght := 0
	in_word := false
	for i := 1; i <= len(s); i++ {
		if s[len(s)-i] != ' ' {
			lenght += 1
			in_word = true
		}

		if s[len(s)-i] == ' ' && in_word {
			return lenght
		}
	}

	return lenght
}
