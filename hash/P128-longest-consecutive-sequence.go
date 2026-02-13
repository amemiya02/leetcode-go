package hash

func longestConsecutive(nums []int) int {

	set := make(map[int]interface{})

	// 放入set 去重
	for _, v := range nums {
		set[v] = struct{}{}
	}

	var res = 0

	// 遍历去重后的set
	for num := range set {
		if _, ok := set[num-1]; !ok {
			// 关键点只有当num-1不存在的时候 才可能是序列的起点
			currentLen := 1
			currentNum := num

			for {
				if _, ok := set[currentNum+1]; ok {
					// 后面有+1递增的数
					currentLen++
					currentNum++
				} else {
					break
				}
			}

			if currentLen > res {
				res = currentLen
			}
		}
	}

	return res
}
