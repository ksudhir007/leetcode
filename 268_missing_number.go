package leetcode

func missingNumber(nums []int) int {
	current_sum := 0
	for i := 0; i < len(nums) ; i = i + 1 {
		current_sum = current_sum + nums[i]
	}
	total_sum := (len(nums) * (len(nums) + 1))/2
	return total_sum - current_sum
}
