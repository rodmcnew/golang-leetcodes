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
    // Then move the left and right pointers inward until we find the match
    for left, right := 0, len(sortedIndicies) - 1; left < right; {
        leftI := sortedIndicies[left]
        rightI := sortedIndicies[right]
        sum := nums[leftI] + nums[rightI]
        if sum > target {
            right--;
        } else if sum < target {
            left++;
        } else {
            return []int{leftI, rightI}
        }
    }

    // The requirements say there is always a solution so panic if its not true.
    panic("No solution found.")
}




//////// Another older solution below ////////




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
    
    // Start with a pointer on each side of the sorted nums       
    // Then move the left and right pointers inward until we find the match
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
    
    // The requirements say there is always a solution so panic if its not true.
    panic("No solution found.")
}
