package api

import (
	"encoding/json"
	"net/http"
)

type Link struct {
	OdataID string `json:"@odata.id"`
}

// We want pretty JSON output, so let's write a function to do this across all handlers
func WritePrettyJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}
