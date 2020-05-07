package main

func isValid(s string) bool {
	size := len(s)
	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++  {
		switch c := s[i];c {
		case '(':
			stack[top] = c + 1
			top++
		case '{','[':
			stack[top] = c + 2
			top++
		case ')','}',']':
			if top > 0 && stack[top-1] == c{
				top--
			} else {
				return false
			}
		}
	}

	return top==0
}

func main(){

	//先進後出
	println(string(40)) // (
	println(string(41)) // )
	println(string(91)) // ]
	println(string(93)) // [
	println(string(123)) // {
	println(string(125)) // }
	println(isValid(""))
}