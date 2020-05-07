package main

import "math"

func mySqrt(x int) int {
	res := x

	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}

func main()  {
	println(mySqrt(8))
	println(math.Sqrt(8))
}
