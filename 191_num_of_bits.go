package leetcode
func hammingWeight(num uint32) int {
	num_bits := 0
	for num != 0 {
		if (num & 1) == 1 {
			num_bits = num_bits + 1
		}
		num = num >> 1
	}
	return num_bits
}
