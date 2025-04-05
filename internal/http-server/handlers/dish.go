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

type DishesRequest struct {
	Limit  int `validate:"omitempty,number,min=0" get:"limit"`
	Offset int `validate:"omitempty,number,min=0" get:"offset"`
}

func (a *API) Dishes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Dishes"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateGetRequest[DishesRequest](w, r, log)
		if !ok {
			return
		}

		dishes, err := a.DishService.Dishes(r.Context(), req.Limit, req.Offset)
		if err != nil {
			log.Error("failed to get dishes", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, dishes)
	}
}
