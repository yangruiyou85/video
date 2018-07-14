package main

func lengthOfNoneRepeatSubStr(s string) int {
	lastCurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastCurred[ch]; ok && lastI >= {
			start = lastI + 1

		}

		if i - start + 1 > maxLength {
			maxLength = i - start + 1

		}
		lastCurred[ch] = i

	}
	return maxLength
}
