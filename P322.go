package leetCode_Go

import (
	"fmt"
	"math"
)

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。
//
//你可以认为每种硬币的数量是无限的。
func test322(){
	arr:=[]int{1,2,5}
	fmt.Print(process(arr,11))
}
func process(nums []int ,amount int)int{
	var dp []int
	for i:=0;i<=amount;i++{
		dp = append(dp, amount+1)
	}
	dp[0] = 0
	for i:=0;i<len(dp);i++{
		for j:=0;j<len(nums);j++{
			if i-nums[j]<0{
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]),float64(dp[i-nums[j]]+1)))
		}
	}
	if dp[amount] == amount+1{
		return -1
	}else {
		return dp[amount]
	}

}

