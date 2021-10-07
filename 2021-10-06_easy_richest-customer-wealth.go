// https://leetcode.com/problems/richest-customer-wealth/

import (
    "fmt"
)

func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for _, amounts := range accounts {
        wealth := 0
        for _, amount := range amounts {
            wealth += amount
        }
        if wealth > maxWealth {
            maxWealth = wealth
        }
    }
    return maxWealth
}
