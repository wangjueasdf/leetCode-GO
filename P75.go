package leetCode_Go

//颜色分类
//给定一个包含红色、白色和蓝色，一共n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
func test75(){
	arr:= []int{2,1,0,2,2,1,0,1}
	sortColors(arr)
}
func sortColors(nums []int){
	if len(nums) <= 1 {
		return
	}
	l, r := partition(nums)
	sortColors(nums[:l])
	sortColors(nums[r:])

}
func partition(nums []int)(int,int){
	cmp := nums[0]
	l,r := 0,len(nums)-1
	for i:=1;i<len(nums)&&i<=r;{
		if nums[i]>cmp{
			nums[i],nums[r] = nums[r],nums[i]
			r--
		} else if nums[i]<cmp{
			nums[i],nums[l] = nums[l],nums[i]
			l++
			i++
		} else{
			i++
		}
	}
	return l,r+1
}