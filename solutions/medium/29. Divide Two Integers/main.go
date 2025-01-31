package main

import "math"

func divide(dividend int, divisor int) int {
	const Max = 2147483647
	const Min = -2147483648

	if dividend == Min && divisor == -1 {
		return Max
	}

	if dividend == 0 {
		return 0
	}

	positive := (dividend < 0) == (divisor < 0)
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	quotient := 0

	if divisor == 1 {
		quotient = dividend
	} else {
		for dividend >= divisor {
			dividend -= divisor
			quotient++
		}
	}

	if positive {
		return quotient
	}
	return -quotient
}
