//go:build ignore

package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	// 结果数组
	ans := make([]int, n-k+1)

	q := []int{}

	for i, v := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] < v {
			q = q[:len(q)-1]
		}

		q = append(q, i)
		// 窗口左边界是 i-k+1 右边界是i
		// 检查队头是否过期
		if q[0] < i-k+1 {
			q = q[1:]
		}
		// 当窗口大小达到k时，记录结果
		if i >= k-1 {
			ans[i-k+1] = nums[q[0]]
		}
	}

	return ans
}
func main() {
	n := 0
	k := 0
	fmt.Scanf("%d %d", &n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	res := maxSlidingWindow(nums, k)
	for i, v := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
