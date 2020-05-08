package main

import "fmt"

func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		n--
		res = string(byte(n%26 + 65)) + res
		n /= 26
	}
	return res
}

func main()  {
	fmt.Println([]byte("A")) //65
	//fmt.Println([]byte("B"))
	//fmt.Println([]byte("Y"))
	//fmt.Println([]byte("Z"))
	//fmt.Println(string(byte(65)))
	fmt.Println((byte(27%26)+'A'))
	fmt.Println(701%26)
	fmt.Println(convertToTitle(701)) //ZY
	fmt.Println(convertToTitle(2)) //B
	fmt.Println(convertToTitle(27)) //AA
	fmt.Println(convertToTitle(26)) //Z
	fmt.Println(convertToTitle(51)) //AY
	fmt.Println(convertToTitle(52)) //AZ
}

//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB
//...

//Input: 701
//Output: "ZY"