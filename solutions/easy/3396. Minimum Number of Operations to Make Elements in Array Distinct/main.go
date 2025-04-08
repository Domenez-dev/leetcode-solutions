func minimumOperations(nums []int) int {
    numsMap := make(map[int]int)
    dis := true
    ops := 0

    for _, num := range nums {
        numsMap[num] += 1
        if numsMap[num] > 1 {
            dis = false
        }
    }

    if dis {
        return ops
    }

    var distinct func(int, []int, map[int]int) int
    distinct = func(ops int, nums []int, numsMap map[int]int) int {
        ops++
        if len(nums) < 3 {
            return ops
        }

        // Remove first 3 elements
        for i := 0; i < 3 && i < len(nums); i++ {
            numsMap[nums[i]] -= 1
        }
        newNums := nums[3:]

        // Check if remaining elements are distinct
        dis := true
        tempMap := make(map[int]int)
        for _, num := range newNums {
            tempMap[num]++
            if tempMap[num] > 1 {
                dis = false
                break
            }
        }

        if dis {
            return ops
        }
        return distinct(ops, newNums, numsMap)
    }

    return distinct(ops, nums, numsMap)
}
