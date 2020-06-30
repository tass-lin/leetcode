package main

func reverseBits(num uint32) uint32 {
	num= (num >> 16) | (num << 16)
	num= ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num= ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num= ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num= ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num

	// for 8 bit binary number abcdefgh, the process is as follow:
	// abcdefgh -> efghabcd -> ghefcdab -> hgfedcba
}

func reverseBit1(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) + (num & 1)
		num = num >> 1
	}
	return res
}



//Example 1:
//
//Input: 00000010100101000001111010011100
//Output: 00111001011110000010100101000000
//Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.

//Example 2:
//
//Input: 11111111111111111111111111111101
//Output: 10111111111111111111111111111111
//Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.