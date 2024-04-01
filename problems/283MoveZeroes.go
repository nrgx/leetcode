package leetcode

// My solution
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	c := 0
	extra := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			extra[c] = nums[i]
			c++
		}
	}
	copy(nums, extra)
}

func moveZeroesBest(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			if left != right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
		right++
	}
}
