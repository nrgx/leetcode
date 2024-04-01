package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	first := nums[0]
	second := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
