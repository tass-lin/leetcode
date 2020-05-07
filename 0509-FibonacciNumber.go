package main

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}

	res := make([]int,N+1)

	res[0], res[1] = 0, 1

	for i := 2; i <= N; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[N-1]+ res[N-2]
}

var fibs = make([]int, 31)

func fib1(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	res := fibs[N]
	if res == 0 {
		res = fib(N-1) + fib(N-2)
		fibs[N] = res
	}
	return res
}

func main()  {
	println(fib(4))
	println(fib1(4))
}