package pcCLub

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"pc_club_server/internal/lib/api/logger/sl"
	"pc_club_server/internal/lib/api/request"
	"pc_club_server/internal/lib/api/response"
)

type PcsRequest struct {
	TypeId int64 `validate:"omitempty,number,min=0" get:"type_id"`
}

func (a *API) Pcs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Pcs"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateGetRequest[PcsRequest](w, r, log)
		if !ok {
			return
		}

		pcs, err := a.PcsService.Pcs(r.Context(), req.TypeId)
		if err != nil {
			log.Error("failed to get pcs", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, pcs)
	}
}
