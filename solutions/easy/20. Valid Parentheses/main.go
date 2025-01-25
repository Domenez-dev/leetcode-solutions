func isValid(s string) bool {
	var open_parenthese []rune
	var last_parenthese rune
	fmt.Println('{', '[')

	if len(s)%2 == 1 {
		return false
	}
	for _, i := range s {
		if i == '{' || i == '(' || i == '[' {
			open_parenthese = append(open_parenthese, i)
		} else if len(open_parenthese) != 0 {
			last_parenthese = open_parenthese[len(open_parenthese)-1]
			fmt.Printf("char: %c\n", i)
			fmt.Printf("open: %c last: %c\n\n", open_parenthese, last_parenthese)
			if (i == '}' && last_parenthese != '{') ||
				(i == ')' && last_parenthese != '(') ||
				(i == ']' && last_parenthese != '[') {
				return false
			}
			open_parenthese = open_parenthese[:len(open_parenthese)-1]
		} else {
			return false
		}
	}
	if len(open_parenthese) == 0 {
		return true
	} else {
		return false
	}
}
