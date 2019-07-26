package leetcode
func singleNumber(nums []int) int {
	x := 0
	for i :=  0 ; i < len(nums) ; i = i + 1 {
		x = x ^ nums[i]
	}
	return x
}
