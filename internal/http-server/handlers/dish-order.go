package pcCLub

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"pc_club_server/internal/lib/api/logger/sl"
	"pc_club_server/internal/lib/api/request"
	"pc_club_server/internal/lib/api/response"
	"pc_club_server/internal/services/order"
)

type OrderDishRequest struct {
	Id    int64 `json:"dish_id" validate:"required,number,min=0"`
	Count int16 `json:"count" validate:"required,number,min=0"`
}

func (a *API) OrderDish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.OrderDish"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		uid := request.MustUID(r)

		req, ok := request.DecodeAndValidateRequest[OrderDishRequest](w, r, log)
		if !ok {
			return
		}

		if err := a.DishOrderService.OrderDish(r.Context(), uid, req.Id, req.Count); err != nil {
			if errors.Is(err, order.ErrNotEnoughMoney) {
				http.Error(w, "not enough money", http.StatusPaymentRequired)
				log.Warn("not enough money", sl.Err(err))
				return
			}

			response.Internal(w)
			log.Error("failed to order dish", sl.Err(err))
			return
		}
	}
}
