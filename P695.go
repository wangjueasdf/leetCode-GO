package leetCode_Go

import "fmt"

//最大岛屿
func test695()  {
	var arr = [][]int {
	{0,0,1,0,0,0,0,1,0,0,0,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,0,0,0,0,0,0,1,1,0,0,0,0}}
	maxAreaOfIsland(arr)
}

func maxAreaOfIsland(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	max := 0
for i:=0;i<row;i++{
	for j:=0;j<col;j++{
		if grid[i][j] == 1{
			max = maxfun(dfs(grid,i,j),max)
		}

	}
}
fmt.Print(max)
return max
}

func dfs( grid [][]int ,row int,col int) (resMax int){
	if row>=len(grid)||row<0||col>=len(grid[0])||col<0||grid[row][col] == 0{
		return 0
	}
	count := 1
	grid[row][col] = 0
	count+= dfs(grid,row+1,col)
	count+= dfs(grid,row-1,col)
	count+= dfs(grid,row,col+1)
	count+= dfs(grid,row,col-1)
	return count
}

func maxfun(a int,b int) (c int){
	if a>b{
		c = a
	}else{
		c = b
	}
	return c
}