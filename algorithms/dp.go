package algorithms

import "weighted-interval/models"

func OptDP(jobs []models.Job, p []int) int {
	n := len(jobs)
	dp := make([]int, n)

	dp[0] = jobs[0].Weight

	for i := 1; i < n; i++ {
		include := jobs[i].Weight

		if p[i] != -1 {
			include += dp[p[i]]
		}

		exclude := dp[i-1]

		if include > exclude {
			dp[i] = include
		} else {
			dp[i] = exclude
		}
	}

	return dp[n-1]
}