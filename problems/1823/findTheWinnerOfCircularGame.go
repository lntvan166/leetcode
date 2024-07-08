/*
https://leetcode.com/problems/find-the-winner-of-the-circular-game/?envType=daily-question&envId=2024-07-08
*/
package exercises1823

func FindTheWinner(n int, k int) int {
	return findTheWinner(n, k)
}

func findTheWinner(n int, k int) int {
	if n == 1 {
		return 1
	}
	isLose := map[int]bool{}
	count := 0
	var result int
	for i := 1; i <= n; i++ {
		isLose[i] = false
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
		}

		winner, isEnd := checkOneStanding(isLose)
		if isEnd {
			result = winner
			break
		}

		i++
		if i > n {
			i = 1
		}
	}

	return result
}

func checkOneStanding(isLose map[int]bool) (int, bool) {
	standingCount := 0
	winner := 0
	for i := 1; i <= len(isLose); i++ {
		if isLose[i] {
			continue
		}

		winner = i
		standingCount++
		if standingCount > 1 {
			return 0, false
		}
	}

	if standingCount == 1 {
		return winner, true
	}
	return 0, false
}
