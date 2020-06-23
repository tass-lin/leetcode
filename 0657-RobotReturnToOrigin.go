package main

import (
	"fmt"
	"strings"
)

func judgeCircle(moves string) bool {
	m := make(map[int32]int, len(moves))

	for _, element := range moves {
		m[element]++
	}

	return  m[85] == m[68] && m[76] == m[82]
}


func judgeCircle1(moves string) bool {
	up := strings.Count(moves, "U")
	down := strings.Count(moves, "D")
	left := strings.Count(moves, "L")
	right := strings.Count(moves, "R")

	return  up == down && left == right
}

func main() {
	fmt.Println([]byte("UDLR"))
	fmt.Println(judgeCircle("UD")) // true
	fmt.Println(judgeCircle("LL")) // false

	fmt.Println(judgeCircle1("UD")) // true
	fmt.Println(judgeCircle1("LL")) // false
}

// There is a robot starting at position (0, 0), the origin, on a 2D plane. Given a sequence of its moves,
// judge if this robot ends up at (0, 0) after it completes its moves.
//
// The move sequence is represented by a string,
// and the character moves[i] represents its ith move. Valid moves are R (right),
// L (left), U (up), and D (down). If the robot returns to the origin after it finishes all of its moves,
// return true. Otherwise, return false.
//
// Note: The way that the robot is "facing" is irrelevant. "R" will always make the robot move to the right once,
// "L" will always make it move left, etc. Also, assume that the magnitude of the robot's movement is the same for each move.
//
// Example 1:
// Input: "UD"
// Output: true
// Explanation: The robot moves up once, and then down once. All moves have the same magnitude,
// so it ended up at the origin where it started. Therefore, we return true.
//
// Example 2:
// Input: "LL"
// Output: false
// Explanation: The robot moves left twice. It ends up two "moves" to the left of the origin.
// We return false because it is not at the origin at the end of its moves.