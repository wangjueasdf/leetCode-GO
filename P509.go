package leetCode_Go

import "fmt"

// 斐波拉契数列
//F(0) = 0，F(1) = 1
//F(n) = F(n - 1) + F(n - 2)，其中 n > 1
func test509(){
	fmt.Println(fba(4))
	fmt.Println(fba(3))
	fmt.Println(fbaDP(4))
	fmt.Println(fbaDP(3))
}

func fba(n int)int{
	if n == 0{
		return 0
	}else if n == 1{
		return 1
	}
	return fba(n-1)+fba(n-2)
}
func fbaDP(n int)int{
	var dp = []int{}
	dp = append(dp, 0)
	dp = append(dp, 1)
	for i:=2;i<=n;i++{
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

