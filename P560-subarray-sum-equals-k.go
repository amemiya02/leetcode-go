//go:build ignore

package main

import "fmt"

func subarraySum(nums []int, k int) int {
	pre := 0
	cnt := make(map[int]int)
	ans := 0
	cnt[0] = 1
	for _, num := range nums {
		pre += num
		if val, ok := cnt[pre-k]; ok {
			ans += val
		}
		cnt[pre]++
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
	k := 0
	fmt.Scanf("%d", &k)

	fmt.Println(subarraySum(nums, k))
}
