package problem0682

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type args struct {
	ops []string
}
var tests = []struct {
	name string
	args args
	want int
}{
	{
		"case1",
		args{[]string{"5", "2", "C", "D", "+"}},
		30,
	},

	{
		"case2",
		args{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}},
		27,
	},
}
func Test_calPoints(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.ops); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tests {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.want, calPoints(tc.args.ops), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			calPoints(tc.args.ops)
		}
	}
}