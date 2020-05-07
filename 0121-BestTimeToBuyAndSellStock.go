package main

func maxProfit(prices []int) int {
	money := 0
	temp := 0
	for i := 1; i < len(prices) ; i++ {
		temp += prices[i] - prices[i-1]

		if temp < 0 {
			temp = 0
		}
		if money < temp {
			money = temp
		}
	}
	return money
}

func main() {
	println(maxProfit([]int{7,1,5,3,6,4}))
}

//Input: [7,1,5,3,6,4]
//Output: 5
//Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//Not 7-1 = 6, as selling price needs to be larger than buying price.
