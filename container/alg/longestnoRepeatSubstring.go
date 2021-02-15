package main

import (
	"fmt"
	"unicode/utf8"
)

func solution(s string) string {
	if len(s) == 0 {
		return ""
	}

	dict := make(map[rune]int)
	start := 0
	var maxLength, left, right int
	for i, ch := range []rune(s) {
		if index, isExist := dict[ch]; isExist && index >= start {
			start = dict[ch] + 1
		}

		if i-start+1 > maxLength {
			left = start
			right = i
			maxLength = i - start + 1
		}

		dict[ch] = i
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
