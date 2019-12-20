package leetcode

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, v := range nums{
		var part = target - v
		if _, ok := m[part]; ok{
			return []int{i, m[part]}
		}
		m[v] = i
	}
	return nil
}

func TestTwoSum(t *testing.T){
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}