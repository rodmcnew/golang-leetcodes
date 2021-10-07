// https://leetcode.com/problems/valid-palindrome/

import (
    "strings"
    "fmt"
)

func isPalindrome(s string) bool {
    // Convert to lower case to make things easier later
    s = strings.ToLower(s)
    
    // Start with a pointer on each side of the string and move inwards.
    // Skip non-alpha-num chars and compare alpha-num chars
    left := 0
    right := len(s) - 1
    for left < right {
        leftChar := s[left]
        if !isAlphaNumeric(leftChar) {
            left++
            continue
        }
        rightChar := s[right]
        if !isAlphaNumeric(rightChar) {
            right--
            continue
        }
        if leftChar != rightChar {
            return false
        }
        left++
        right--
    }
    return true
}

// This purposefully only works on lower case letters
func isAlphaNumeric (s byte) bool {
    return (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')
}
