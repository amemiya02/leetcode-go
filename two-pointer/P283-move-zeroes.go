package twopointer

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
