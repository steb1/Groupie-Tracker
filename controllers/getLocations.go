package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getLocations(s string, w http.ResponseWriter) ResponseLocations {
	response1, err := http.Get(s)
	if err != nil {
		error(w, "500")
		w.WriteHeader(500)
	}
	responseData1, err1 := io.ReadAll(response1.Body)
	if err1 != nil {
		error(w, "500")
		w.WriteHeader(500)
	}
	var responseObjectLocations ResponseLocations
	json.Unmarshal(responseData1, &responseObjectLocations)

	return responseObjectLocations
}
