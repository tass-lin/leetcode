package main

import (
	"math"
	"strings"
	"fmt"
)

func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string){
	fmt.Println(s)
	s = strings.TrimSpace(s)
	fmt.Println(s)
	if s == "" {
		return
	}

	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs = ""
		return
	}

	for i,b := range abs {
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}
	return
}

func convert(sign int, absStr string) int {
	abs := 0
	for _, b := range absStr{
		abs = abs * 10 + int(b-'0')

		switch {
		case sign == 1 && abs > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && sign*abs < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * abs
}

func main(){
	string := " "
	ans := myAtoi(string)
	fmt.Println(ans)
}