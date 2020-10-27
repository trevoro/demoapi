package render

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/trevoro/demoapi/errors"
)

type errorEnvelope struct {
	Error error `json:"error"`
}

func mustIndent(data interface{}) []byte {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(fmt.Errorf("error encoding json: %w", err))
	}
	return b
}

func Error(w http.ResponseWriter, err error) {
	var buffer []byte
	var status int

	// This is where we can put specific error type handling for global error
	// context / normalizing. For example, if we ever run into a valdiation
	// error, then we can intercept that and write it in a clean format. Or if
	// we run into errors with persistence we can normalize those too. For now
	// there is just one default error case. TODO: For validator errors we
	// should be adding some details to the error envelope.
	switch v := err.(type) {
	case validator.ValidationErrors:
		buffer = mustIndent(errorEnvelope{errors.ErrBadRequest})
		status = errors.ErrBadRequest.Status
	default:
		log.Printf("there was an unknown error T[%T]: %v", v, v)
		buffer = mustIndent(v)
		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, e := w.Write(buffer)
	if e != nil {
		log.Println("error writing response:", e)
	}
}
