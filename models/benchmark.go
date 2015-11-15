// Package models contains all the models within the analysis project.
package models

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// Benchmark represents the data that is associated with performing a benchmark.
// It includes a name, which represents the benchmark function, Operations,
// which represents the number of operations that were executed to find the
// Performance number.
type Benchmark struct {
	gorm.Model
	Name        string
	Operations  int64
	Performance float64
	Project     Project
}

var matchExp = regexp.MustCompile(`(?P<name>^.*)\s+(?P<operations>\d+)\s+(?P<performance>\b([0-9]+\.[0-9])|(\d+)?)\s`)
var expNames = matchExp.SubexpNames()

// BenchmarkFromLine will take a string line that is in an expected benchmark
// format and create a struct from it.
// BenchmarkTest1                20000000                89.4 ns/op
// BenchmarkTest2                20000000                28.5 ns/op
func BenchmarkFromLine(line string) (*Benchmark, error) {
	matches := matchExp.FindStringSubmatch(line)
	if len(matches) >= len(expNames) {
		data := getNamedMatches(matches)

		performance, err := strconv.ParseFloat(data["performance"], 64)
		if err != nil {
			return nil, err
		}

		operations, err := strconv.ParseInt(data["operations"], 10, 64)
		if err != nil {
			return nil, err
		}

		return &Benchmark{
			Name:        data["name"],
			Operations:  operations,
			Performance: performance,
		}, nil
	}

	return nil, nil
}

func getNamedMatches(matches []string) map[string]string {
	data := make(map[string]string)
	for i, m := range expNames {
		if m == "" {
			continue
		}

		data[m] = strings.TrimSpace(matches[i])
	}

	return data
}
