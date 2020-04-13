package leetCode_Go

import (
	"fmt"
	"strings"
)

func test1160()  {
	words := []string  {"cat","bt","hat","tree"}
	chars := "atach"
	fmt.Print(countCharacters(words,chars))
}

func countCharacters(words []string, chars string) int {
	count:=0
outLoop :for i:=0;i<len(words);i++{
	for _,bt := range words[i]{
		if strings.Count(words[i],string(bt))>strings.Count(chars,string(bt)){
			continue outLoop
		}
	}
	count+=len(words[i])
	}
return count
}