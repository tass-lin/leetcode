package main

import "fmt"

// 兩個人，我先撿石頭，最多三顆
// 如為 5 ， 取 1 後 對方取幾都會輸 ， 如為6 則 2 ， 如為 7 則 3
// 發現 4 與 8 都會因為無法留超過三顆給對方，同理推測4倍數都會輸

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(canWinNim(4)) // false
	fmt.Println(canWinNim(5)) // true
	fmt.Println(canWinNim(6)) // true
	fmt.Println(canWinNim(7)) // true
	fmt.Println(canWinNim(8)) // false
}

//You are playing the following Nim Game with your friend: There is a heap of stones on the table, each time one of you take turns to remove 1 to 3 stones. The one who removes the last stone will be the winner. You will take the first turn to remove the stones.
//Both of you are very clever and have optimal strategies for the game. Write a function to determine whether you can win the game given the number of stones in the heap.
//
//Example:
//Input: 4
//Output: false
//Explanation: If there are 4 stones in the heap, then you will never win the game;
//No matter 1, 2, or 3 stones you remove, the last stone will always be removed by your friend.