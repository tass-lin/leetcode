package main

func maxProfit(prices []int) int {
	temp := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1]  {
			temp += prices[i] - prices[i-1]
		}
	}
	return temp
}

func main() {
	println(maxProfit([]int{7,1,5,3,6,4})) //7
	println(maxProfit([]int{1,2,3,4,5})) //4
	println(maxProfit([]int{6,1,3,2,4,7})) //7
	println(maxProfit([]int{7,6,4,3,1})) //0
	println(maxProfit([]int{2,1,4})) //3
}

//Input: [7,1,5,3,6,4]
//Output: 7
//Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
//Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

//Input: [1,2,3,4,5]
//Output: 4
//Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
//engaging multiple transactions at the same time. You must sell before buying again.