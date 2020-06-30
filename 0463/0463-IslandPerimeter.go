package main

import "fmt"

// N*4 - M*2

func islandPerimeter(grid [][]int) int {
	m, n:= len(grid), len(grid[0])
	counts, neighbors := 0, 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				counts++
				if  i < m - 1 {
					if grid[i + 1][j] == 1 {
						neighbors++
					}
				}
				if j < n -1 {
					if grid[i][j + 1] == 1 {
						neighbors++
					}
				}
			}
		}
	}

	return 4 * counts - 2 * neighbors
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func islandPerimeter1(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4
			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
					res--
				}
			}
		}
	}

	return res
}

func main() {
	arr := [][]int{[]int{0, 1, 0, 0},[]int{1, 1, 1, 0},[]int{0, 1, 0, 0},[]int{1, 1, 0, 0}}
	fmt.Println(islandPerimeter(arr)) // 16
	fmt.Println(islandPerimeter1(arr)) // 16
	fmt.Println(islandPerimeter2(arr)) // 16

	// [[1]] 4
	// [[1,1],[1,1]] 8
}

// You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.
// Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
// and there is exactly one island (i.e., one or more connected land cells).
// The island doesn't have "lakes" (water inside that isn't connected to the water around the island).
// One cell is a square with side length 1. The grid is rectangular,
// width and height don't exceed 100. Determine the perimeter of the island.



// Example:
// Input:
// [[0,1,0,0],
// [1,1,1,0],
// [0,1,0,0],
// [1,1,0,0]]
//
// Output: 16
//
// Explanation: The perimeter is the 16 yellow stripes in the image below: