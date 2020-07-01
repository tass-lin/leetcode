package main

import "testing"

type args struct {
	s string
}
var tests = []struct {
	name string
	args args
	want bool
}{
	{
		"case1",
		args{"aba"},
		true,
	},
	{
		"case2",
		args{"abca"},
		true,
	},
	{
		"case3",
		args{"tebbem"},
		false,
	},
	{
		"case4",
		args{"abc"},
		false,
	},
	{
		"case5",
		args{"abcbac"},
		true,
	},
	{
		"case6",
		args{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"},
		true,
	},
}
func Test_validPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validPalindrome1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			validPalindrome(tc.args.s)
		}
	}
}

func Benchmark_fn1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			validPalindrome1(tc.args.s)
		}
	}
}