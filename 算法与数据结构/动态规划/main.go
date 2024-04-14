package main

//对于动态规划问题，我将拆解为如下五步曲，这五步都搞清楚了，才能说把动态规划真的掌握了！
//
//确定dp数组（dp table）以及下标的含义
//确定递推公式
//dp数组如何初始化
//确定遍历顺序
//举例推导dp数组

//子序列问题：
//单个序列：  以dp[i]即是以i为结尾的.
//
//
//两个序列： dp[i][j]： 以i-1,j-1为结尾的
//	连续：  不用处理不等于的情况
//  不连续： 需要处理不等于的情况   num1[i] != num[j]   dp[i-1]dp[j]    dp[i]dp[j-1]
