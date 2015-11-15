// Package controllers defines a set of controllers and handlers that are used
// for the web server.
package controllers

import (
	"net/http"

	"github.com/jelmersnoeck/analysis/models"
	"github.com/jelmersnoeck/analysis/utils"
)

func IndexProjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`foo`))
}

func CreateProjectsHandler(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}
	err := utils.Unmarshal(r, &project)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
