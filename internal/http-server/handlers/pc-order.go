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

type OrderPcRequest struct {
	PcId      int64 `json:"pc_id" validate:"required,number,min=0"`
	HourCount int16 `json:"hour_count" validate:"required,number,min=0"`
}

func (a *API) OrderPc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.OrderPc"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		uid := request.MustUID(r)

		req, ok := request.DecodeAndValidateRequest[OrderPcRequest](w, r, log)
		if !ok {
			return
		}

		code, err := a.PcOrderService.OrderPc(r.Context(), uid, req.PcId, req.HourCount)
		if err != nil {
			log.Warn("failed to create order", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, code)
	}
}
