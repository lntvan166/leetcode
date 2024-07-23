// https://leetcode.com/problems/sort-array-by-increasing-frequency/description/?envType=daily-question&envId=2024-07-23
package exercises1636

import "sort"

func frequencySort(nums []int) []int {
	mapFrequency := map[int]int{}
	arrayFrequency := []int{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		mapFrequency[nums[i]]++
	}
	for k := range mapFrequency {
		arrayFrequency = append(arrayFrequency, k)
	}
	sort.Slice(arrayFrequency, func(i, j int) bool {
		if mapFrequency[arrayFrequency[i]] == mapFrequency[arrayFrequency[j]] {
			return arrayFrequency[i] > arrayFrequency[j]
		}
		return mapFrequency[arrayFrequency[i]] < mapFrequency[arrayFrequency[j]]
	})
	for i := 0; i < len(arrayFrequency); i++ {
		for j := 0; j < mapFrequency[arrayFrequency[i]]; j++ {
			result = append(result, arrayFrequency[i])
		}
	}
	return result
}
