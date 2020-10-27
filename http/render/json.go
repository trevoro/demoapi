package render

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// This should probably be a method on a struct where we store options and
// create the writers on New() time instead. That way there are no if
// statements in here but we could have mutliple behaviours in the same
// application.
func JSON(w http.ResponseWriter, status int, data interface{}) {
	buff, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(fmt.Errorf("error encoding data: %w", err))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err = w.Write(buff)
	if err != nil {
		log.Println("error writing response:", err)
	}
}
