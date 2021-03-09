package main

import (
	"testing"
	"unicode/utf8"
)

func TestSubstring(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"ababa", 2},
		{"pwwkew", 3},
		{"abcdef", 6},

		// Edge cases
		{"bbbbb", 1},
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"慕课网慕课", 3},
		{"黑化黑灰化肥灰挥发发灰黑讳为灰黑花会回飞", 8},
	}

	for _, tt := range tests {
		if actual := solution(tt.s); utf8.RuneCountInString(actual) != tt.ans {
			t.Errorf("noRepeatSubStringsolution(%s); got %s(%d); expect %d", tt.s, actual, utf8.RuneCountInString(actual), tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化黑灰化肥灰挥发发灰黑讳为灰黑花会回飞"
	ans := 8
	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s)= %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actualS := solution(s)
		actual := utf8.RuneCountInString(actualS)
		if actual != ans {
			b.Errorf("noRepeatSubStringsolution(%s); got %s(%d); expect %d", s, actualS, actual, ans)
		}
	}
}
