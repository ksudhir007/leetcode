package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	char_count := make(map[rune]int)
	for _, val := range(s){
		char_count[val] = char_count[val] + 1
	}

	for _, val := range(t) {
		char_count[val] = char_count[val] - 1
	}

	for _, v := range(char_count){
		if v != 0 {
			return false
		}
	}
	return true
}
