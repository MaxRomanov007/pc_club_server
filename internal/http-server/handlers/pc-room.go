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

type PcRoomRequest struct {
	RoomId int64 `get:"room-id,true" validate:"required,numeric,min=1"`
}

type PcRoomsRequest struct {
	PcTypeID int64 `get:"type_id" validate:"required,numeric,min=1"`
}

func (a *API) PcRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pcClub.PcRooms"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateGetRequest[PcRoomsRequest](w, r, log)
		if !ok {
			return
		}

		pcRooms, err := a.PcRoomService.PcRooms(r.Context(), req.PcTypeID)
		if err != nil {
			log.Error("failed to get pc room", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, pcRooms)
	}
}
