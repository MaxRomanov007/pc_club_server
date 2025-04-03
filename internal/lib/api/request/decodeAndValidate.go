package request

import (
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"pc_club_server/internal/lib/api/logger/sl"
	"pc_club_server/internal/lib/api/response"
	"pc_club_server/internal/lib/request/urlGet"
)

func DecodeAndValidateRequest[T any](
	w http.ResponseWriter,
	r *http.Request,
	log *slog.Logger,
) (T, bool) {
	var req T
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Error("failed to decode request body", sl.Err(err))
		response.Internal(w)
		return req, false
	}

	if !ValidateRequest[T](w, req, log) {
		return req, false
	}

	return req, true
}

func DecodeAndValidateGetRequest[T any](
	w http.ResponseWriter,
	r *http.Request,
	log *slog.Logger,
) (T, bool) {
	var req T
	if err := urlGet.Decode(r, &req); err != nil {
		log.Error("failed to decode get request", sl.Err(err))
		response.Internal(w)
		return req, false
	}

	if !ValidateRequest[T](w, req, log) {
		return req, false
	}

	return req, true
}
