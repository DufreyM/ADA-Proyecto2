package algorithms

import "weighted-interval/models"

func OptRecursive(jobs []models.Job, p []int, i int) int {
	if i < 0 {
		return 0
	}

	include := jobs[i].Weight
	if p[i] != -1 {
		include += OptRecursive(jobs, p, p[i])
	}

	exclude := OptRecursive(jobs, p, i-1)

	if include > exclude {
		return include
	}
	return exclude
}