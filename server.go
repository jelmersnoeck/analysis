package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("benchmark")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	myExp := regexp.MustCompile(`(?P<name>^.*)\s+(?P<operations>\d+)\s+(?P<performance>\b([0-9]+\.[0-9])|(\d+)?)\s`)
	names := myExp.SubexpNames()
	benchmarks := make([]Benchmark, 0)

	for scanner.Scan() {
		matches := myExp.FindStringSubmatch(scanner.Text())
		if len(matches) >= len(names) {
			data := make(map[string]string)
			for i, m := range names {
				if m == "" {
					continue
				}

				data[m] = strings.TrimSpace(matches[i])
			}

			p, _ := strconv.ParseFloat(data["performance"], 64)
			i, _ := strconv.ParseInt(data["operations"], 10, 64)
			b := Benchmark{
				Name:        data["name"],
				Operations:  i,
				Performance: p,
			}

			benchmarks = append(benchmarks, b)
		}
	}

	fmt.Println(benchmarks)
}

type Benchmark struct {
	Name        string
	Operations  int64
	Performance float64
}
