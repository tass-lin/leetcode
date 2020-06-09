package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if r * c != m * n {
		return nums
	}
	res := make([][]int, r)

	for i := 0; i < r * c; i++ {
		res[i/c] = append(res[i/c], nums[i/n][i%n])
	}
	return res
}

func main() {
	nums := [][]int{[]int{1, 2}, []int{3, 4}}
	fmt.Printf("%v", matrixReshape(nums, 1, 4)) // [[1 2 3 4]]
	fmt.Printf("%v", matrixReshape(nums, 2, 4)) // [[1 2] [3 4]]

	nums1 := [][]int{[]int{1, 2, 3, 4}}
	fmt.Printf("%v", matrixReshape(nums1, 2, 2)) // [[1 2] [3 4]]
}

// In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.
//
// You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.
//
// The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.
//
// If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.
//
// Example 1:
// Input:
// nums =
// [[1,2],
// [3,4]]
// r = 1, c = 4
// Output:
// [[1,2,3,4]]
// Explanation:
// The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.

// Example 2:
// Input:
// nums =
// [[1,2],
// [3,4]]
// r = 2, c = 4
// Output:
// [[1,2],
// [3,4]]
// Explanation:
// There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.