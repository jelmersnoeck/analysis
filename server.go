package main

import (
	"fmt"
	"net/http"

	"golang.org/x/tools/benchmark/parse"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", CreateHandler).Methods("POST")

	http.ListenAndServe(":5000", r)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	f, _, err := r.FormFile("benchmark")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	set, err := parse.ParseSet(f)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(set)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`OK`))
}
