// https://leetcode.com/problems/median-of-two-sorted-arrays/
package exercises4

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays2(nums1, nums2)
}

// my solution O(nlogn)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a := append(nums1, nums2...)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	n := len(a)
	i := int(n / 2)
	if n%2 == 1 {
		return float64(a[i])
	} else {
		return float64(a[i-1]+a[i]) / 2
	}
}

// leetcode solution O(log(min(m,n)))
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays2(nums2, nums1)
	}
	x, y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		partitionX := low + (high-low)/2
		partitionY := (x+y+1)/2 - partitionX
		maxLeftX, minRightX := 0, 0
		if partitionX != 0 {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX != x {
			minRightX = nums1[partitionX]
		}
		maxLeftY, minRightY := 0, 0
		if partitionY != 0 {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY != y {
			minRightY = nums2[partitionY]
		}
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0
}
