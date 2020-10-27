package http

import (
	"net/http"

	"github.com/go-playground/form"
	"github.com/go-playground/validator/v10"
)

var (
	decoder  *form.Decoder
	validate *validator.Validate
)

func init() {
	decoder = form.NewDecoder()
	validate = validator.New()
}

func bind(r *http.Request, v interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := decoder.Decode(v, r.Form); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}
