// https://leetcode.com/problems/lucky-numbers-in-a-matrix/description/?envType=daily-question&envId=2024-07-19
package exercises1380

func luckyNumbers(matrix [][]int) []int {
	minMap := map[int]bool{}
	maxMap := map[int]bool{}
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		min := matrix[i][0]
		for j := 0; j < m; j++ {
			if min > matrix[i][j] {
				min = matrix[i][j]
			}
		}
		minMap[min] = true
	}
	for i := 0; i < m; i++ {
		max := matrix[0][i]
		for j := 0; j < n; j++ {
			if max < matrix[j][i] {
				max = matrix[j][i]
			}
		}
		maxMap[max] = true
	}
	result := []int{}
	for k := range minMap {
		if _, ok := maxMap[k]; ok {
			result = append(result, k)
		}
	}
	return result
}
