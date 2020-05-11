package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	arrByte := []byte(s)
	res, size := 0, len(arrByte)

	for i := 0; i < size; i++ {
		temp := int(arrByte[i] - 64)
		res = res * 26 + temp
	}
	return res
}

func main() {
	//fmt.Println([]byte("A")) //65
	//fmt.Println(string(byte(65))) // A
	//fmt.Println([]byte("AA")) // [65 65]
	//fmt.Println([]byte("ZY")) // [90 89]

	fmt.Println(titleToNumber("A")) //1
	fmt.Println(titleToNumber("AA")) //27
	fmt.Println(titleToNumber("AB")) //28
	fmt.Println(titleToNumber("ZY")) //701
	fmt.Println(titleToNumber("AAA")) //703
	fmt.Println(titleToNumber("AAB")) //704

}

//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
