//go:build ignore

package main

import "fmt"

func findAnagrams(s string, p string) []int {
	l1 := len(s)
	l2 := len(p)
	if l1 < l2 {
		return []int{}
	}
	ans := []int{}

	cnt2 := [26]int{}
	cnt1 := [26]int{}
	for i := 0; i < l2; i++ {
		cnt2[p[i]-'a']++
		cnt1[s[i]-'a']++
	}

	if cnt1 == cnt2 {
		ans = append(ans, 0)
	}

	for i := l2; i < l1; i++ {

		cnt1[s[i]-'a']++
		cnt1[s[i-l2]-'a']--

		if cnt1 == cnt2 {
			ans = append(ans, i-l2+1)
		}
	}
	return ans

}

func main() {
	s := ""
	p := ""
	fmt.Scanf("%s %s", &s, &p)
	res := findAnagrams(s, p)
	for i, v := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
