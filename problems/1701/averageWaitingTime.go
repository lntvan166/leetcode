package exercises1701

func AverageWaitingTime(customers [][]int) float64 {
	return averageWaitingTime(customers)
}

func averageWaitingTime(customers [][]int) float64 {
	var waitTime, currentTime float64
	for _, customer := range customers {
		if currentTime < float64(customer[0]) {
			currentTime = float64(customer[0])
		}
		currentTime += float64(customer[1])
		waitTime += currentTime - float64(customer[0])
	}
	return waitTime / float64(len(customers))
}
