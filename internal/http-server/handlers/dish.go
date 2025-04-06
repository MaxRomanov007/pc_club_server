package pcCLub

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"pc_club_server/internal/lib/api/logger/sl"
	"pc_club_server/internal/lib/api/request"
	"pc_club_server/internal/lib/api/response"
	"pc_club_server/internal/storage/mssql"
)

type DishesRequest struct {
	Limit  int `validate:"omitempty,number,min=0" get:"limit"`
	Offset int `validate:"omitempty,number,min=0" get:"offset"`
}

type DishRequest struct {
	Id int64 `validate:"required,number,min=1" get:"id, true"`
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

func (a *API) Dish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Dish"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateGetRequest[DishRequest](w, r, log)
		if !ok {
			return
		}

		dish, err := a.DishService.Dish(r.Context(), req.Id)
		if err != nil {
			if errors.Is(err, mssql.ErrNotFound) {
				log.Warn("not found", sl.Err(err))
				response.NotFound(w)
			}
			log.Error("failed to get pc type", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, dish)
	}
}
