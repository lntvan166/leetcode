package exercises1823

/*
There are n friends that are playing a game. The friends are sitting in a circle and are numbered from 1 to n in clockwise order. More formally, moving clockwise from the ith friend brings you to the (i+1)th friend for 1 <= i < n, and moving clockwise from the nth friend brings you to the 1st friend.

The rules of the game are as follows:

1. Start at the 1st friend.
2. Count the next k friends in the clockwise direction including the friend you started at. The counting wraps around the circle and may count some friends more than once.
3. The last friend you counted leaves the circle and loses the game.
4. If there is still more than one friend in the circle, go back to step 2 starting from the friend immediately clockwise of the friend who just lost and repeat.
5. Else, the last friend in the circle wins the game.
Given the number of friends, n, and an integer k, return the winner of the game.
*/
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
