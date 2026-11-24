package main

import "fmt"

// SubsetSum reports whether some subset of nums sums exactly to target.
// One boolean row over achievable sums, swept backwards per element so
// each number is used at most once.
func SubsetSum(nums []int, target int) bool {
	if target < 0 {
		return false
	}
	dp := make([]bool, target+1)
	dp[0] = true // the empty subset
	for _, v := range nums {
		if v < 0 || v > target {
			continue // negative values unsupported; too-large values useless
		}
		for s := target; s >= v; s-- {
			if dp[s-v] {
				dp[s] = true
			}
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(SubsetSum([]int{3, 34, 4, 12, 5, 2}, 9))
	fmt.Println(SubsetSum([]int{3, 34, 4, 12, 5, 2}, 30))
}
