//go:build ignore

package main

import "fmt"

func moveZeroes(nums []int) {

	p := 0

	for q := 0; q < len(nums); q++ {
		if nums[q] != 0 {
			nums[p] = nums[q]
			p++
		}
	}

	for ; p < len(nums); p++ {
		nums[p] = 0
	}

}

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	moveZeroes(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(nums[i])
	}
	fmt.Println()
}
