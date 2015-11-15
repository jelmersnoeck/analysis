// Package utils provides several helper functions to use throughout the code
// base
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Unmarshal(r *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
