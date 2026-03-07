//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[target-nums[i]]; ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	target := 0
	fmt.Scanf("%d", &target)
	res := twoSum(nums, target)

	fmt.Println(strconv.Itoa(res[0]) + " " + strconv.Itoa(res[1]))
}
