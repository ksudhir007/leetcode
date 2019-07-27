package leetcode
import (
	"strconv"
)
func isPalindrome(x int) bool {
	num_str := ""

	if x < 0 {
		return false
	}
	for x > 0 {
		digit := x % 10
		num_str = num_str + strconv.Itoa(digit)
		x = (x - digit)/10
	}

	length := len(num_str) - 1
	for  i,j := 0, length; i <= j; i, j = i + 1,j - 1 {
		if num_str[i] != num_str[j] {
			return false
		}
	}

	return true
}