package algorithms

import "weighted-interval/models"

func ComputeP(jobs []models.Job) []int {
	n := len(jobs)
	p := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = -1
		for j := i - 1; j >= 0; j-- {
			if jobs[j].End <= jobs[i].Start {
				p[i] = j
				break
			}
		}
	}

	return p
}