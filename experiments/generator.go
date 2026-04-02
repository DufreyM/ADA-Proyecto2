package experiments

import (
	"math/rand"
	"weighted-interval/models"
	"sort"
)

func GenerateJobs(n int) []models.Job {
	jobs := make([]models.Job, n)

	current := 0

	for i := 0; i < n; i++ {
		start := current
		duration := rand.Intn(50) + 1
		end := start + duration
		weight := rand.Intn(100) + 1

		jobs[i] = models.Job{
			Start:  start,
			End:    end,
			Weight: weight,
		}

		// forzar solapamiento parcial
		current += rand.Intn(20)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].End < jobs[j].End
	})

	return jobs
}