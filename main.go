package main

import (
	"fmt"
	"time"
	"encoding/csv"
	"os"
	"weighted-interval/algorithms"
	"weighted-interval/experiments"
)

func main() {

	// Crear carpeta si no existe
	os.MkdirAll("results", os.ModePerm)

	file, err := os.Create("results/data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"n", "recursive", "dp"})

const RUNS = 100

for n := 1; n <= 30; n++ {

	jobs := experiments.GenerateJobs(n)
	p := algorithms.ComputeP(jobs)

	var totalRec int64 = 0
	var totalDP int64 = 0

	for r := 0; r < RUNS; r++ {

		// Recursivo
		startRec := time.Now()
		resultRec := 0
		for i := 0; i < 20; i++ {
			resultRec += algorithms.OptRecursive(jobs, p, n-1)
		}
		totalRec += time.Since(startRec).Nanoseconds()

		// DP
		startDP := time.Now()
		resultDP := 0
		for i := 0; i < 2000; i++ {
			resultDP += algorithms.OptDP(jobs, p)
		}
		totalDP += time.Since(startDP).Nanoseconds()

		// evitar optimización
		if resultRec == -1 || resultDP == -1 {
			fmt.Println("Impossible")
		}
	}

	avgRec := totalRec / RUNS
	avgDP := totalDP / RUNS

	writer.Write([]string{
		fmt.Sprintf("%d", n),
		fmt.Sprintf("%d", avgRec),
		fmt.Sprintf("%d", avgDP),
	})
}
	// IMPORTANTE: verificar errores del writer
	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
}