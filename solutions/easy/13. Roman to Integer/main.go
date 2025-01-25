func checkNext(s string, symbols map[rune]int, a, b rune, i int) (int, int) {
	fmt.Printf("check func- %v: %c \n", i, s[i])
	if i+1 < len(s) {
		if rune(s[i+1]) == b {
			return symbols[b] - symbols[a], i + 1
		} else if rune(s[i+1]) == a {
			result := 0
			result, i = checkNext(s, symbols, a, b, i+1)
			return result + symbols[a], i
		}
	}
	return symbols[a], i
}

func romanToInt(s string) int {
	symbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	result := 0

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v: %c \n", i, s[i])
		switch s[i] {
		case 'I':
			if i+1 < len(s) {
				if rune(s[i+1]) == 'V' {
					result, i = checkNext(s, symbols, 'I', 'V', i)
				} else {
					result, i = checkNext(s, symbols, 'I', 'X', i)
				}
			} else {
				result = 1
			}
			total += result

		case 'X':
			if i+1 < len(s) {
				if rune(s[i+1]) == 'L' {
					result, i = checkNext(s, symbols, 'X', 'L', i)
				} else {
					result, i = checkNext(s, symbols, 'X', 'C', i)
				}
			} else {
				result = 10
			}
			total += result

		case 'C':
			if i+1 < len(s) {
				if rune(s[i+1]) == 'D' {
					result, i = checkNext(s, symbols, 'C', 'D', i)
				} else {
					result, i = checkNext(s, symbols, 'C', 'M', i)
				}
			} else {
				result = 100
			}
			total += result

		default:
			total += symbols[rune(s[i])]
		}
	}

	return total
}
