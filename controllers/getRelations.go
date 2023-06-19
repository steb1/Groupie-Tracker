package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getRelations(s string, w http.ResponseWriter) ResponseRelation {
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
	var responseObjectRelations ResponseRelation
	json.Unmarshal(responseData1, &responseObjectRelations)

	return responseObjectRelations
}