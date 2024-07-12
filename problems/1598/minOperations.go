// https://leetcode.com/problems/crawler-log-folder/?envType=daily-question&envId=2024-07-10
package exercises1598

// my solution 2ms
func minOperations(logs []string) int {
	stack := []string{}
	for _, log := range logs {
		if log == "../" && len(stack) != 0 {
			stack = stack[:len(stack)-1]
			continue
		}
		if log != "./" && log != "../" {
			stack = append(stack, log)
		}
	}

	return len(stack)
}

// leetcode solution 0ms
func minOperationsLeetCode(logs []string) int {
	var target int
	for _, log := range logs {
		if target < 0 {
			target = 0
		}
		if log[0] == 46 {
			if log[1] == 46 {
				target--
			}
		} else {
			target++
		}
	}
	if target < 0 {
		return 0
	}
	return target
}
