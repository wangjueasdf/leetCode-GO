package leetCode_Go

import "fmt"

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
var res [][]int
func testP46()  {
	a :=[]int{1,2,3}
	fmt.Print(permute(a))
}


func permute(nums []int) [][]int {
	res = [][]int{}
	backtrack(nums,[]int{},len(nums))
	return res
}
func backtrack(nums[]int,nowArr[]int, length int){
	if len(nums) == 0{
		p:=make([]int,len(nowArr))
		copy(p,nowArr)
		res = append(res, p)
	}
	for i:=0;i<length;i++{
		cur:=nums[i]
		nowArr = append(nowArr, nums[i])
		nums = append(nums[:i],nums[i+1:]...)
		backtrack(nums,nowArr,len(nums))
		nowArr = append(nowArr[:len(nowArr)-1])
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
	}

}
