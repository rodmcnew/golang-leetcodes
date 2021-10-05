// https://leetcode.com/problems/two-sum/

import (
    "fmt"
    "sort"
)

func twoSum(nums []int, target int) []int {
    // Create a slice of the num indicies sorted by num value.
    // This gives us a sorted view of the nums while retaining their original indicies.
    sortedIndicies := make([]int, len(nums))
    for i := range nums {
        sortedIndicies[i] = i
    }
    sort.Slice(sortedIndicies, func(a, b int) bool {
        return nums[sortedIndicies[a]] < nums[sortedIndicies[b]]
    })

    // Start with a pointer on each side of the sorted nums
    left := 0
    right := len(sortedIndicies) - 1
        
    // Move the left and right pointers inward until we find the match
    for left < right {
        leftI := sortedIndicies[left]
        rightI := sortedIndicies[right]
        
        leftVal := nums[leftI]
        rightVal := nums[rightI]
        
        sum := leftVal + rightVal
        
        if sum > target {
            right--;
        } else if sum < target {
            left++;
        } else {
            return []int{leftI, rightI}
        }
    }

    panic("No solution found.")
}
