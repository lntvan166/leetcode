// https://leetcode.com/problems/two-sum/description/
package exercises1

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
