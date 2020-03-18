package leetCode_Go

import (
	"fmt"
	"math"
)

func test836()  {
	var rec1  = []int{4,0,6,6}
	var rec2 = []int {-5,-3,4,2}
	fmt.Println(isRectangleOverlap(rec1,rec2))
}
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	var c1 = [2]float64 {float64((rec1[2]-rec1[0])/2) + float64(rec1[0]),float64((rec1[3]-rec1[1])/2) + float64(rec1[1])}
	var c2 = [2]float64 {float64((rec2[2]-rec2[0])/2) + float64(rec2[0]),float64((rec2[3]-rec2[1])/2) + float64(rec2[1])}
	var sumY = c1[1]-float64(rec1[1]) +c2[1]-float64(rec2[1])
	var sumX = c1[0]-float64(rec1[0]) +c2[0]-float64(rec2[0])
	if math.Abs(c1[0]-c2[0])< (sumX) &&math.Abs(c1[1]-c2[1])< (sumY){
		return true
	}else{
		return false
	}
}

func shadow(rec1 []int, rec2 []int) bool {
	return  !(rec1[0]>=rec2[2]||rec1[2]<=rec2[0]) && !(rec1[1]>=rec2[3]||rec1[3]<=rec2[1])
}