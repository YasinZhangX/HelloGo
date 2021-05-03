package main

import (
	"fmt"
	"unicode/utf8"
)

var lastOccured = make([]int, 0xffff)

func solution(s string) string {
	if len(s) == 0 {
		return ""
	}

	// lastOccured := make(map[rune]int)
	// lastOccured := make([]int, 0xffff)
	for i := range lastOccured {
		lastOccured[i] = -1
	}

	start := 0
	var maxLength, left, right int
	for i, ch := range []rune(s) {
		if index := lastOccured[ch]; index != -1 && index >= start {
			start = lastOccured[ch] + 1
		}

		if i-start+1 > maxLength {
			left = start
			right = i
			maxLength = i - start + 1
		}

		lastOccured[ch] = i
	}

	return string([]rune(s)[left : right+1])
}

func test(s string) {
	result := solution(s)
	fmt.Println(result, utf8.RuneCountInString(result))
}

func main() {
	test("ababa")
	test("bbbbb")
	test("")
	test("pwwkew")
	test("b")
	test("abcdef")
	test("慕课网慕课")
	test("阿巴阿巴阿")
}
