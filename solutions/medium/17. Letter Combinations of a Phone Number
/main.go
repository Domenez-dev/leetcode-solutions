func letterCombinations(digits string) []string {
	combos := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}

	i := 0 // len(digits)
	return comboFinder(i, digits, combos)
}

func comboFinder(i int, digits string, combos map[string][]string) []string {
	if i == len(digits)-1 {
		return combos[fmt.Sprintf("%c", digits[i])]
	}

	result := []string{}
	for _, h := range combos[fmt.Sprintf("%c", digits[i])] {
		for _, j := range comboFinder(i+1, digits, combos) {
			result = append(result, h+j)
		}
	}
	return result
}
