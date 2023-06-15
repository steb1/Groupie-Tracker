package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getRelations(s string) ResponseRelation {
	response1, _ := http.Get(s)
	responseData1, _ := io.ReadAll(response1.Body)
	var responseObjectRelations ResponseRelation
	json.Unmarshal(responseData1, &responseObjectRelations)

	return responseObjectRelations
}