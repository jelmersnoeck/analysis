package models_test

import (
	"testing"

	"github.com/jelmersnoeck/analysis/models"
)

func TestFromCorrectLine(t *testing.T) {
	line := "BenchmarkTest1                20000000                89.4 ns/op"
	b, _ := models.BenchmarkFromLine(line)

	if b.Name != "BenchmarkTest1" {
		t.Errorf("Name does not match BenchmarkTest1")
	}

	if b.Operations != 20000000 {
		t.Errorf("Operations does not match 20000000")
	}

	if b.Performance != 89.4 {
		t.Errorf("Performance does not match 89.4")
	}
}

func TestFromInvalidLine(t *testing.T) {
	line := "Benchmark Is Invalid"

	b, _ := models.BenchmarkFromLine(line)

	if b != nil {
		t.Errorf("Benchmark should be nil")
	}
}

func TestFromIncorrectLine(t *testing.T) {
	line := "BenchmarkTest1                String                89.4 ns/op"
	b, _ := models.BenchmarkFromLine(line)

	if b != nil {
		t.Errorf("Benchmark should be nil")
	}
}
