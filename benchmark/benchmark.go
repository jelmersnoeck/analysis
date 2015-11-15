// Package benchmark handles benchmark comparison.
package benchmark

import (
	"regexp"
	"strconv"
	"strings"
)

// Benchmark represents the data that is associated with performing a benchmark.
// It includes a name, which represents the benchmark function, Operations,
// which represents the number of operations that were executed to find the
// Performance number.
type Benchmark struct {
	Name        string
	Operations  int64
	Performance float64
}

var matchExp *regexp.Regexp = regexp.MustCompile(`(?P<name>^.*)\s+(?P<operations>\d+)\s+(?P<performance>\b([0-9]+\.[0-9])|(\d+)?)\s`)
var expNames []string = matchExp.SubexpNames()

func FromLine(line string) (*Benchmark, error) {
	matches := matchExp.FindStringSubmatch(line)

	if len(matches) >= len(expNames) {
		data := make(map[string]string)
		for i, m := range expNames {
			if m == "" {
				continue
			}

			data[m] = strings.TrimSpace(matches[i])
		}

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
