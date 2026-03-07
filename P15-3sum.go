//go:build ignore

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	var ans = [][]int{}

	if n < 3 {
		return ans
	}
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			// 当前比0大，后面的会更大 直接结束循环
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			// 去重
			continue
		}

		target := -nums[i]

		// 双指针查找
		l := i + 1
		r := n - 1

		for l < r {
			temp := nums[l] + nums[r]
			if temp == target {
				// 找到一组解
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				// 先移动指针
				l++
				r--
				// 去重
				for nums[l] == nums[l-1] && l < r {
					l++
				}
				for nums[r] == nums[r+1] && l < r {
					r--
				}
			} else if temp > target {
				r--
			} else {
				l++
			}
		}
	}

	return ans
}

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	ans := threeSum(nums)
	for i := 0; i < len(ans); i++ {
		fmt.Println(strconv.Itoa(ans[i][0]) + " " + strconv.Itoa(ans[i][1]) + " " + strconv.Itoa(ans[i][2]))
	}
}
