package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jelmersnoeck/analysis/models"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Benchmark{})

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
