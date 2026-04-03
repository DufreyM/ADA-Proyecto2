package experiments

import (
	"time"
)

func MeasureTimeRecursive(f func() int) (int, int64) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start).Nanoseconds()
	return result, elapsed
} 