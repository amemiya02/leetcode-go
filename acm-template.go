package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

// 初始化scanner，在每个题目文件的main函数里调用一次
func InitScanner() {
	scanner = bufio.NewScanner(os.Stdin)
}

// 读取下一个整数
func NextInt() int {
	val := 0
	fmt.Scanf("%d", &val)
	return val
}

// 读取下一个字符串（以空格/换行分隔）
func Next() string {
	val := ""
	fmt.Scanf("%s", &val)
	return val
}

// 读取一整行字符串（包含空格）
func NextLine() string {
	scanner.Scan()
	return scanner.Text()
}

// 读取一行中的多个字符串
func NextStrs() []string {
	var res []string
	str := NextLine()
	split := strings.Split(str, " ")
	for _, s := range split {
		res = append(res, s)
	}
	return res
}

func NextInts() []int {
	strs := NextStrs()
	n := len(strs)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		v, _ := strconv.Atoi(strs[i])
		res[i] = v
	}
	return res
}

func main() {
	InitScanner()
	nums := NextInts()
	fmt.Println(nums)
}
