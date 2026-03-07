# LeetCode Hot 100 Golang Solutions

## 哈希

- 1. 两数之和
  - [题目](https://leetcode-cn.com/problems/two-sum/)
  - [解法](./P1-two-sum.go)
- 49. 字母异位词分组
  - [题目](https://leetcode-cn.com/problems/group-anagrams/)
  - [解法](./P49-group-anagrams.go)
- 128. 最长连续序列
  - [题目](https://leetcode-cn.com/problems/longest-consecutive-sequence)
  - [解法](./P128-longest-consecutive-sequence.go)

## 双指针

- 283. 移动零
  - [题目](https://leetcode-cn.com/problems/move-zeroes/)
  - [解法](./P283-move-zeroes.go)
- 11. 盛最多水的容器
  - [题目](https://leetcode-cn.com/problems/container-with-most-water/)
  - [解法](./P11-container-with-most-water.go)
- 15. 三数之和
  - [题目](https://leetcode-cn.com/problems/3sum/)
  - [解法](./P15-3sum.go)
- 42. 接雨水
  - [题目](https://leetcode-cn.com/problems/trapping-rain-water/)
  - [解法](./P42-trapping-rain-water.go)

## 滑动窗口

- 3. 无重复字符的最长子串
  - [题目](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
  - [解法](./P3-longest-substring-without-repeating-characters.go)
- 438. 找到字符串中所有字母异位词
  - [题目](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)
  - [解法](./P438-find-all-anagrams-in-a-string.go)

## 子串

- 560. 和为 K 的子数组
  - [题目](https://leetcode-cn.com/problems/subarray-sum-equals-k/)
  - [解法](./P560-subarray-sum-equals-k.go)
- 239. 滑动窗口最大值
  - [题目](https://leetcode-cn.com/problems/sliding-window-maximum/)
  - [解法](./P239-sliding-window-maximum.go)


## bufio的输入写法

```go
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 读取一行输入（比如："1 0 3 0 5"）
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// 按空格分割字符串
	strSlice := strings.Fields(line)

	// string切片转int切片
	nums := make([]int, len(strSlice))
	for i, s := range strSlice {
		// 笔试中直接忽略错误（题目保证输入合法）
		nums[i], _ = strconv.Atoi(s)
	}

	// int切片转回string
	var strResult string
	for i, num := range nums {
		if i > 0 {
			strResult += " "
		}
		strResult += strconv.Itoa(num)
	}
}
```