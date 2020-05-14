package main

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	arrSecret := []rune(secret)
	arrGuess := []rune(guess)
	mSecret := make(map[rune]int,len(arrSecret))
	mGuess := make(map[rune]int,len(arrGuess))
	var bull, cows int
	for index, element := range arrSecret {
		if element == arrGuess[index] {
			bull++
		} else {
			mSecret[element] +=1
			mGuess[arrGuess[index]] +=1
		}
	}
	for index, element := range mSecret {
		if mGuess[index] <= element  {
			cows+= mGuess[index]
		} else {
			cows+= element
		}
	}

	return strconv.Itoa(bull)+"A"+strconv.Itoa(cows)+"B"
}
func getHint1(secret string, guess string) string {
	var bull, cow int
	var countS, countG [10]int

	size := len(secret)
	for i := 0; i < size; i++ {
		ns := int(secret[i] - '0')
		ng := int(guess[i] - '0')

		if ns == ng {
			bull++
		} else {
			if countG[ns] > 0 {
				cow++
				countG[ns]--
			} else {
				countS[ns]++
			}
			if countS[ng] > 0 {
				cow++
				countS[ng]--
			} else {
				countG[ng]++
			}
		}
	}

	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}

func main() {
	fmt.Println(getHint("1807","7810")) // 1A3B 107 710
	fmt.Println(getHint("1123","0111")) // 1A1B 123 011
	fmt.Println(getHint("1122","0001")) // 0A1B

	fmt.Println(getHint1("1807","7810")) // 1A3B 107 710
	fmt.Println(getHint1("1123","0111")) // 1A1B 123 011
	fmt.Println(getHint1("1122","0001")) // 0A1B
}

//You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint
//that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows").
//Your friend will use successive guesses and hints to eventually derive the secret number.
//Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows.
//Please note that both secret number and friend's guess may contain duplicate digits.
//
//Example 1:
//Input: secret = "1807", guess = "7810"
//Output: "1A3B"
//Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.

//Example 2:
//Input: secret = "1123", guess = "0111"
//Output: "1A1B"
//Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.