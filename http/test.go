package http

import (
	"net/http"

	"github.com/trevoro/demoapi/http/render"
)

func (s *Server) handleTest() http.HandlerFunc {
	type request struct {
		Name  string `form:"name" validate:"required"`
		Email string `form:"email" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var p request
		if err := bind(r, &p); err != nil {
			render.Error(w, err)
			return
		}
		render.Record(w, http.StatusOK, p, nil)
	}
}
