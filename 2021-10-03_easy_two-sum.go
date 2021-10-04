// https://leetcode.com/problems/two-sum/

import (
	"fmt"
    "sort"
)
	
type num struct {
    index int
    value int
}

func twoSum(nums []int, target int) []int {
    
    // Get the nums sorted while remembering their original indices
    var sortedNums = make([]num,len(nums))
    for index, value := range nums {
        sortedNums[index].index = index
        sortedNums[index].value = value
    }
    sort.Slice(sortedNums, func(a int, b int) bool {return sortedNums[a].value < sortedNums[b].value})
    // fmt.Print("sorted Nums", sortedNums)
    
    // Move the left and right pointers inward until we find the match
    var left int = 0
    var right = len(nums) - 1
    for left < right {
        var sum = sortedNums[left].value + sortedNums[right].value
        // fmt.Print("in loop [", left, sortedNums[left].value, "] + [", right, sortedNums[right].value, "] = ", sum, "\n");
        if sum > target {
            right--;
        } else if sum < target {
            left++;
        } else {
            return []int{sortedNums[left].index,sortedNums[right].index}
        }
    }
    
    // This should never happen but is needed for typing
    return []int{}
}
