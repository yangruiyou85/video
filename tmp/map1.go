package main

import (
	"fmt"
)

func main() {

	fmt.Println(FindStr("abcdabc"))
	fmt.Println(FindStr("我爱你北京我。"))

}

//寻找最长不含有重复字符的子串

func FindStr(s string) int {

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i

	}

	return maxLength

}
