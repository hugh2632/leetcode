package leetcode

import (
	"math"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sum = len(nums1) + len(nums2)
	//是否是偶数个
	var isOdd = sum&1 == 0
	//中位数左一位或正中位
	var leftIndex = (sum - 1) / 2
	//需要移除的个数
	var totalRemoved = leftIndex
	for totalRemoved > 0 {
		var once = (totalRemoved + 1) / 2
		var index1, index2 = 0, 0
		var len1, len2 = len(nums1), len(nums2)
		if len1 > len2 {
			len1, len2 = len2,len1
			nums1, nums2 = nums2, nums1
		}
		if len1 == 0{
			break
		}
		if len1 >= once{
			index1 = once - 1
		}else{
			index1 = len1 - 1
		}
		index2 = once - 1
		if nums1[index1] > nums2[index2]{
			nums2 = nums2[index2 + 1 :]
			totalRemoved -= index2 + 1
		}else{
			nums1 = nums1[index1 + 1 :]
			totalRemoved -= index1 + 1
		}
	}
	var len1, len2 = len(nums1), len(nums2)
	if len1 > len2 {
		nums1, nums2 = nums2, nums1
		len1,len2 = len2, len1
	}
	if len1 == 0 {
		if len2 == 0 {
			return 0
		}else {
			if isOdd {
				return float64(nums2[totalRemoved] + nums2[totalRemoved + 1]) / 2
			}else{
				return float64(nums2[totalRemoved])
			}
		}
	}else{
		if len2 == 0 {
			if isOdd {
				return float64(nums1[0] + nums1[1]) / 2
			}else{
				return float64(nums1[0])
			}
		}else{
			if isOdd{
				if nums1[0] < nums2[0] {
					if len(nums1) > 1 {
						return (math.Min(float64(nums1[1]), float64(nums2[0])) + float64(nums1[0])) / 2
					}else{
						return float64(nums1[0] + nums2[0]) / 2
					}
				} else{
					if len(nums2) > 1 {
						return (math.Min(float64(nums2[1]), float64(nums1[0])) + float64(nums2[0])) / 2
					}else{
						return float64(nums2[0] + nums1[0]) / 2
					}
				}
			}else{
				return math.Min(float64(nums1[0]),float64(nums2[0]))
			}
		}
	}


}

func TestFindMedianSortedArrays(t *testing.T) {
	var nums1 = []int{1,3}
	var nums2 = []int{2}
	t.Log(findMedianSortedArrays(nums1,nums2))
}
