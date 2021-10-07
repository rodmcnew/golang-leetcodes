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




//////// Concurrency isn't really needed for this problem but below is a solution that uses concurrency anyways ////////



import (
    "fmt"
    "sync"
)

func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    
    var wg sync.WaitGroup
    var mu sync.Mutex

    for _, amounts := range accounts {
        wealth := 0
        amounts := amounts
        wg.Add(1)
        go func () {
            defer wg.Done()
            for _, amount := range amounts {
                wealth += amount
            }
            mu.Lock()
            if wealth > maxWealth {
                maxWealth = wealth
            }
            mu.Unlock()
        }()
    }
    
    wg.Wait();
    
    return maxWealth
}
