package leetcode

import (
"fmt"
"testing"
)

func lengthOfLongestSubstring(s string) int {
	var store = make([]rune, 0)
	var maxLength = 0
	for _, v := range s {
		for i, _ := range store {
			if store[i] == v {
				store = store[i+1:]
				break
			}
		}
		store = append(store, v)
		if len(store) > maxLength {
			maxLength = len(store)
		}
	}
	return maxLength
}

func TestLengthOfLongest(t *testing.T) {
	var res = lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}
