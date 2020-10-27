package render

import (
	"net/http"
)

// Record makes it easy for us to push results that return paired results
// directly into this method without an error check. It is just sugar.
func Record(w http.ResponseWriter, status int, data interface{}, err error) {
	if err != nil {
		Error(w, err)
		return
	}
	JSON(w, status, data)
}
