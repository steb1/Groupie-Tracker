package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getDates(s string) ResponseDates {
	response, _ := http.Get(s)
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectDates ResponseDates
	json.Unmarshal(responseData, &responseObjectDates)

	return responseObjectDates
}