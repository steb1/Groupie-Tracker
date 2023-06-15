package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func getArtistes(s string) ResponseArtist {
	response, _ := http.Get(s)
	responseData, _ := io.ReadAll(response.Body)
	var responseObjectArtist ResponseArtist
	json.Unmarshal(responseData, &responseObjectArtist)

	return responseObjectArtist
}
