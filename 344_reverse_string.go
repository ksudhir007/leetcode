package leetcode

func reverseString(s []byte)  {
	lStr := len(s) - 1
	for i,j := 0,lStr; i < lStr && i <= j ; i, j = i+1, j-1 {
		t := s[i]
		s[i] = s[j]
		s[j] = t
	}

}
