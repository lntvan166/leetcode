// https://leetcode.com/problems/sort-the-people/?envType=daily-question&envId=2024-07-22
package exercises2418

func sortPeople(names []string, heights []int) []string {
	mapPeople := map[int]string{}
	var result []string
	for i := range names {
		mapPeople[heights[i]] = names[i]
	}

	sortedHeights := quickSortStart(heights)
	for _, h := range sortedHeights {
		result = append([]string{mapPeople[h]}, result...)
	}

	return result
}

func quickSort(arr []int, low, high int) []int {

	if low < high {

		var p int

		arr, p = partition(arr, low, high)

		arr = quickSort(arr, low, p-1)

		arr = quickSort(arr, p+1, high)

	}

	return arr

}

func quickSortStart(arr []int) []int {

	return quickSort(arr, 0, len(arr)-1)

}

func partition(arr []int, low, high int) ([]int, int) {

	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {

		if arr[j] < pivot {

			arr[i], arr[j] = arr[j], arr[i]

			i++

		}

	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i

}
