package main

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		"case1",
		args{[]int{1,3,5,4,7}},
		3,
	},
	{
		"case1",
		args{[]int{2,2,2,2,2}},
		1,
	},
	{
		"case1",
		args{[]int{1,3,5,7}},
		4,
	},
}
func Test_findLengthOfLCIS(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests{
			findLengthOfLCIS(tc.args.nums)
		}
	}
}