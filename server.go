package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jelmersnoeck/analysis/models"
)

func main() {
	file, err := os.Open("test-benchmark")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	benchmarks := make([]*models.Benchmark, 0)

	for scanner.Scan() {
		b, err := models.BenchmarkFromLine(scanner.Text())

		if err != nil {
			fmt.Println(err)
			break
		}

		if b != nil {
			benchmarks = append(benchmarks, b)
		}
	}

	fmt.Println(benchmarks)
}
