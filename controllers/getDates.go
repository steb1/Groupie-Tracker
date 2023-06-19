package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getDates(s string, w http.ResponseWriter) ResponseDates {
	response, err := http.Get(s)
	if err != nil {
		error(w, "500")
		w.WriteHeader(500)
	}
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectDates ResponseDates
	json.Unmarshal(responseData, &responseObjectDates)

	return responseObjectDates
}