import "math"

//dp solution

func coinChange(coins []int, amount int) int {
    dp := []int{0}
    
    for i := 0; i < amount; i++{
        dp = append(dp, math.MaxInt64)
    }

    for i := 1; i <= amount; i++{
        for j := 0; j < len(coins); j++{
            if i-coins[j] >= 0 && 
            dp[i-coins[j]] != math.MaxInt64{
                dp[i] = min(dp[i], 1+dp[i-coins[j]])
            }
        }
    }
    // fmt.Println(dp)
    if dp[amount] == math.MaxInt64{
        return -1
    }
    return dp[amount]
}

func min(a int, b int) int{
    if a < b{
        return a
    }
    return b
}

// solution1 with memo, not dp
func coinChange(coins []int, amount int) int {
    mp := make(map[int]int)
    ans := dp(coins, amount, mp)
    return ans
}

func dp(coins []int, amount int, memo map[int]int) int{
    if amount < 0{
        return -1
    }
    if amount == 0{
        return 0
    }
    if _, ok := memo[amount]; ok{
        return memo[amount]
    }

    count := math.MaxInt64
    for i := 0; i < len(coins); i++{
        step := dp(coins, amount-coins[i], memo)
        if step == -1{
            continue
        }
        count = min(count, step+1)
    }
    if count == math.MaxInt64{
        memo[amount] = -1 
    }else{
        memo[amount] = count
    }
    return memo[amount]
}

func min(a int, b int) int{
    if a  > b{
        return b
    }
    return a
}