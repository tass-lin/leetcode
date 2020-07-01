package main

import (
	"fmt"
	"math"
)

//123
func reverse(x int) int {

	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0

	for x > 0 {
		temp := x % 10
		res = res * 10 + temp
		x = x / 10
	}

	res = sign * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}


	return res
}
func main(){
	num := 1
	fmt.Println(reverse(num))
}