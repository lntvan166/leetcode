// https://leetcode.com/problems/find-the-winner-of-the-circular-game/?envType=daily-question&envId=2024-07-08
package exercises1823

func FindTheWinner(n int, k int) int {
	return findTheWinner(n, k)
}

func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}
	isLose := map[int]bool{}
	player := map[int]bool{}
	count := 0
	var result int
	for i := 1; i <= n; i++ {
		isLose[i] = false
		player[i] = true
	}

	i := 1
	for {
		if isLose[i] {
			i++
			if i > n {
				i = 1
			}
			continue
		}

		count++
		if count == k {
			count = 0
			isLose[i] = true
			delete(player, i)
		}

		if len(player) == 1 {
			break
		}

		i++
		if i > n {
			i = 1
		}
	}

	for k := range player {
		result = k
	}

	return result
}
