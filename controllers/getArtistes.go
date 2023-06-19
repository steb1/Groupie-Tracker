package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getArtistes(s string, w http.ResponseWriter) ResponseArtist {
	response, err := http.Get(s)
	if err != nil {
		error(w, "500")
		w.WriteHeader(500)
	}
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	return responseObjectArtist
}
