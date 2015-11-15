package models

import "github.com/jinzhu/gorm"

// Project represents a project which will be analysed.
type Project struct {
	gorm.Model
	Name       string `json:"name"`
	Benchmarks []Benchmark
}
