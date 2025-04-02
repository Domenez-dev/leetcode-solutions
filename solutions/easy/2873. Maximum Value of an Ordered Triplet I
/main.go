package main

func maximumTripletValue(nums []int) int64 {
    var n int64 = 0
    for i:=0; i<len(nums)-2; i++{
        for j:=i+1; j<len(nums)-1; j++{
            if nums[j] < nums[i] {
                for k:=j+1; k<len(nums); k++ {
                    n = max(n, int64((nums[i] - nums[j]) * nums[k]))
                }
            }
        }
    }
    return n
}

func max(x, y int64) int64{
    if x > y {
        return x
    }
    return y
}
