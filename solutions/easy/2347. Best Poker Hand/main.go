func bestHand(ranks []int, suits []byte) string {
	flush := func(suits []byte) bool {
		for _, i := range suits {
			if i != suits[0] {
				return false
			}
		}
		return true
	}

	parity := func(ranks []int) (bool, bool) {
		var count int
		pair := false

		for i := 0; i < 3; i++ {
			count = 0
			for j := i + 1; j < len(ranks); j++ {
				if ranks[j] == ranks[i] {
					count += 1
					if count == 2 {
						return true, true
					}
					if count == 1 {
						pair = true
					}
				}
			}
		}
		return false, pair
	}

	three, pair := parity(ranks)
	switch true {
	case flush(suits):
		return "Flush"
	case three:
		return "Three of a Kind"
	case pair:
		return "Pair"
	default:
		return "High Card"
	}
}
