package main

import (
	"fmt"
	"github.com/astaxie/beego"
)
//1071 
func main()  {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Print(gcdOfStrings(str1,str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1{
		return ""
	}
	return beego.Substr(str1,0,gcd(len(str1),len(str2)))
}

func gcd(a int,b int) int{
	if(a%b == 0){
		return b
	}else{
		return gcd(b,a%b)
	}
}