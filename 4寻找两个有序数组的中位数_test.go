package leetcode

import (
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
		var len1 = len(nums1)
		var len2 = len(nums2)
		if len1 > len2 {
			nums1, nums2 = nums2, nums1
			len1, len2 = len2, len1
		}
		//1.判断len1 + len2是奇数还是偶数
		var isOdd = (len1 + len2) & 1 != len1 + len2
		//2.得到要获得的中位数的左侧一个数
		
		//3.通过折半比大小去掉小的数，留下大的数继续比，直到取到左边数

		//4.通过判断是奇还是偶来取最后的值


	return 0.9
}

func TestFindMedianSortedArrays(t *testing.T) {

}
