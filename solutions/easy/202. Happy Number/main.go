package main

import (
    "strconv"
    "math"
)

func isHappy(n int) bool {
    mymap := map[int]bool{}
    var sum int
    var num string
    for {
        if mymap[n] == true {
            return false
        }
        num = strconv.Itoa(n)
        sum = 0
        for _, i := range num {
            digit, err := strconv.Atoi(string(i))
            if err != nil {
                return false
            }
            sum += int(math.Pow(float64(digit), 2))
        }
        if sum == 1 {
            return true
        }
        mymap[n] = true
        n = sum
    }
}
