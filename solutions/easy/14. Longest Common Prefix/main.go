func longestCommonPrefix(strs []string) string {
	var prefix string
	first_string := strs[0]
	for i := 0; i < len(first_string); i++ {
		for _, j := range strs[1:] {
			if len(j) <= i || j[i] != first_string[i] {
				return prefix
			}
		}
		prefix += string(first_string[i])
	}
	return prefix
}
