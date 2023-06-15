package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getLocations(s string) ResponseLocations {
	response1, _ := http.Get(s)
	responseData1, _ := io.ReadAll(response1.Body)
	var responseObjectLocations ResponseLocations
	json.Unmarshal(responseData1, &responseObjectLocations)

	return responseObjectLocations
}
