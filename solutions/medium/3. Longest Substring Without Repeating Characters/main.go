func lengthOfLongestSubstring(s string) int {
	subString := ""
	prevString := ""
	for i := 0; i < len(s); i++ {
		subString = noRepeatCheck(s[i:])
		if len(subString) > len(prevString) {
			prevString = subString
		}
	}
	return len(prevString)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func noRepeatCheck(s string) string {
	var used []rune
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(used); j++ {
			if used[j] == rune(s[i]) {
				return s[:i]
			}
		}
		used = append(used, rune(s[i]))
	}
	return s
}
