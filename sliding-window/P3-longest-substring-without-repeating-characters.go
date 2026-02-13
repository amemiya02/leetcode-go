package slidingwindow

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	left, right := 0, 0
	ans := 0
	countMap := make(map[byte]int)
	for right < n {
		// 当前字符出现过 要从左侧收缩
		for {
			if v, ok := countMap[s[right]]; ok {
				delete(countMap, s[left])
				left = 
			} else {
				break
			}
		}
		countMap[s[right]] = struct{}{}
		if right-left+1 > ans {
			ans = right - left + 1
		}
		right++

	}
	return ans
}
