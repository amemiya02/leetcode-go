package hash

func groupAnagrams(strs []string) [][]string {

	m := make(map[[26]int][]string)

	for _, str := range strs {
		cnt := [26]int{}
		for i := 0; i < len(str); i++ {
			cnt[str[i]-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}

	res := make([][]string, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
