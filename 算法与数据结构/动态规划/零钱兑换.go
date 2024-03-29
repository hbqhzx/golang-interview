package main

import "math"

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
//你可以认为每种硬币的数量是无限的。
//示例 1：
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//示例 2：
//
//输入：coins = [2], amount = 3
//输出：-1
//示例 3：
//
//输入：coins = [1], amount = 0
//输出：0

//for j := coins[i]; j <= amount; j++：
//这种方式是“完全背包”问题的解决方式，用于处理每种硬币可以使用多次的情况。它首先固定一种硬币，然后试图用这种硬币组合出可能的所有金额。
//
//for j := amount; j >= coins[i]; j--：
//这种方式是“0-1背包”问题的解决方式，用于处理每种硬币只能使用一次的情况。它首先固定一种硬币，然后尝试接近总金额。

// 完全背包
func coinChange(coins []int, amount int) int {
	//定义dp[j], 表示，拼成j金额所需要的最少硬币数量dp[j]
	dp := make([]int, amount+1)

	//初始化：首先凑足总金额为0所需钱币的个数一定是0，那么dp[0] = 0;
	//其他下标对应的数值呢？
	//考虑到递推公式的特性，dp[j]必须初始化为一个最大的数，否则就会在min(dp[j - coins[i]] + 1, dp[j])比较的过程中被初始值覆盖。
	//
	//所以下标非0的元素都是应该是最大值。
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	//推断dp, dp[j] = min(dp[j-coins[i]]+1,dp[j])
	// j的值每次都是coins[i]开始的，因为j要大于coins[i], 下面的dp[j-conins[i]] 计算才有意义
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			//检查 dp[j] 是否可能溢出
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j-coins[i]]+1, dp[j])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
