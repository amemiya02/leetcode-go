package twopointer

func trap(height []int) int {

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(height)

	left := make([]int, n)
	right := make([]int, n)

	// left[i] 0到i的最大值
	// right[i] i到n-1的最大值
	left[0] = height[0]
	right[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += (min(left[i], right[i]) - height[i])
	}
	return ans
}
