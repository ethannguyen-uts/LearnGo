package main

import (
	"fmt"
)

func moveZeroes(nums []int) []int {
    var nonZeroPointer int = 0;
    for i, _ := range nums {
        fmt.Printf("%v", i)
        for nonZeroPointer <= i || (nonZeroPointer < len(nums) && nums[nonZeroPointer] == 0) {
            nonZeroPointer += 1
        }

        if (nonZeroPointer >= len(nums)){
            break
        }

        if (nums[i] == 0) {
            nums[i], nums[nonZeroPointer] = nums[nonZeroPointer], nums[i]
        }
    }
    return nums
}