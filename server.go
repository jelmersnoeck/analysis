package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jelmersnoeck/analysis/controllers"
	"github.com/jelmersnoeck/analysis/models"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Benchmark{})

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.IndexProjectsHandler).Methods("GET")
	r.HandleFunc("/projects", controllers.CreateProjectsHandler).
		Methods("POST").Headers("Content-Type", "application/json")

	http.ListenAndServe(":8000", r)
}
