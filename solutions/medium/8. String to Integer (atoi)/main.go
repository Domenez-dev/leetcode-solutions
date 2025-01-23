func myAtoi(s string) int {
	digits := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	positive := true
	result := 0

	if len(s) > 0 {
		for len(s) > 0 && digits[rune(s[0])] == 0 && rune(s[0]) != '0' {
			if s[0] == ' ' {
				s = s[1:]
				continue

			} else if rune(s[0]) == '-' || rune(s[0]) == '+' {
				if rune(s[0]) == '-' {
					positive = false
				}
				s = s[1:]
				break

			} else {
				return result
			}
		}

		for len(s) > 0 && (digits[rune(s[0])] != 0 || rune(s[0]) == '0') {
			result = result*10 + digits[rune(s[0])]
			if result > math.MaxInt32 {
				break
			}
			s = s[1:]
		}
	}

	if positive {
		if result > math.MaxInt32 || result < math.MinInt32 {
			return math.MaxInt32
		}
	}

	if !positive {
		result = -result
		if result < math.MinInt32 || result > math.MaxInt32 {
			return math.MinInt32
		}
	}

	return result
}
