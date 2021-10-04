// https://leetcode.com/problems/two-sum/

import (
    "fmt"
    "sort"
)

func twoSum(nums []int, target int) []int {
    
    // Get the nums into a data structure that we can sort while remembering their indicies
    sortedNums := make([]struct{index, value int}, len(nums))
    for index, value := range nums {
        sortedNums[index].index = index
        sortedNums[index].value = value
    }
    
    // Sort the nums
    sort.Slice(sortedNums, func(a, b int) bool {return sortedNums[a].value < sortedNums[b].value})
    
    // Move the left and right pointers inward until we find the match
    left := 0
    right := len(nums) - 1
    for left < right {
        sum := sortedNums[left].value + sortedNums[right].value
        if sum > target {
            right--;
        } else if sum < target {
            left++;
        } else {
            return []int{sortedNums[left].index, sortedNums[right].index}
        }
    }
    
    // This should never happen but is needed for typing
    return []int{}
}
