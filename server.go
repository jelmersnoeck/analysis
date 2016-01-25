package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/tools/benchmark/parse"
)

func main() {
	file, err := os.Open("test-benchmark")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	benchmarks := make([]*parse.Benchmark, 0)

	for scanner.Scan() {
		b, err := parse.ParseLine(scanner.Text())

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
