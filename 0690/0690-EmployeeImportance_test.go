package problem0690

import "testing"

/*[[1,2,[2]], [2,3,[]]]
2*/

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}

	employee1 := Employee{1, 5, []int{2, 3}}
	employee2 := Employee{2, 3, []int{}}
	employee3 := Employee{3, 3, []int{}}
	var employees []*Employee
	employees = append(employees,&employee1)
	employees = append(employees,&employee2)
	employees = append(employees,&employee3)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
		"case1",
		args{employees, 1},
		11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}