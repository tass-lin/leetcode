package main

import "math"

func longestCommonPrefix(strs []string) string {
	short := shortTest(strs)
	for i, r := range short {
		for j := 0; j < len(strs); j++{
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}
	return short
}

func shortTest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]

	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}
	return res
}

func main(){

	//var strs []string
	//strs[0] = "flower"
	//strs[1] = "flow"
	//strs[2] = "flight"
	strs := []string{"flower","floyer","flight"}

	println(longestCommonPrefix(strs))

}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := math.MaxInt32
	minStr := ""
	//先找到最短的字符串
	for _, v := range strs {
		l := len(v)
		if l < minLen {
			minLen = l
			minStr = v
		}
	}
	long := len(minStr)
	for _, v := range strs {
		for i := 0; i < minLen; i++ {
			if minStr[i:i+1] != v[i:i+1] {
				long = int(math.Min(float64(long), float64(i)))
			}
		}
	}
	return minStr[:long]
}