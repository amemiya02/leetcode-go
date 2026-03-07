//go:build ignore

package main

import "fmt"

func maxArea(height []int) int {

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

	l, r := 0, len(height)-1
	ans := (r - l) * min(height[l], height[r])

	for l < r {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		temp := (r - l) * min(height[l], height[r])
		ans = max(ans, temp)
	}

	return ans

}

func main() {

	n := 0
	fmt.Scanf("%d", &n)
	height := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &height[i])
	}
	fmt.Println(maxArea(height))
}
