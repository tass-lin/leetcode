package main

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= num - 1
	}
	return res
}

func main()  {
	var num uint32
	num = 00000000000000000000000000001011
	println(hammingWeight(num)) //3
	num = 00000000000000000000000010000000
	println(hammingWeight(num)) //1
}
//Example 1:
//
//Input: 00000000000000000000000000001011
//Output: 3
//Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.

//Example 2:
//
//Input: 00000000000000000000000010000000
//Output: 1
//Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.

//Example 3:
//
//Input: 11111111111111111111111111111101
//Output: 31
//Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.
