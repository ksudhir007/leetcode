package leetcode
import (
	"fmt"
)
func fizzBuzz(n int) []string {
	output := make([]string, 0)
	for i := 1; i <= n; i = i + 1  {
		if i % 3 == 0 && i %5 == 0 {
			output = append(output, "FizzBuzz")
		} else if i % 5 == 0 {
			output = append(output, "Buzz")
		} else if i % 3 == 0 {
			output = append(output, "Fizz")
		} else {
			s := fmt.Sprintf("%d", i)
			output = append(output, s)
		}
	}
	return output
}
