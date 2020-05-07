package main


func strStr(haystack string, needle string) int {
	haystackLen , needleLen := len(haystack), len(needle)

	for i := 0; i <= haystackLen-needleLen ; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

func main()  {
	println(strStr("hello","ll")) //haystackLen = 5
	println(strStr("aaaaa","bba")) //haystackLen = 5 , needleLen = 3
	println(strStr("","a"))
	println(strStr("",""))
}