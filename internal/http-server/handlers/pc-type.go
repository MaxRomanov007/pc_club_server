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

type PcTypesRequest struct {
	Limit  int `validate:"omitempty,number,min=0" get:"limit"`
	Offset int `validate:"omitempty,number,min=0" get:"offset"`
}

type PcTypeRequest struct {
	TypeId int64 `validate:"required,number,min=1" get:"type-id, true"`
}

func (a *API) PcTypes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pcClub.pcType.PcTypes"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateGetRequest[PcTypesRequest](w, r, log)
		if !ok {
			return
		}

		pcTypes, err := a.PcTypeService.PcTypes(r.Context(), req.Limit, req.Offset)
		if err != nil {
			log.Error("failed to get pc types", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, pcTypes)
	}
}

func (a *API) PcType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.pcClub.pcType.PcType"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateGetRequest[PcTypeRequest](w, r, log)
		if !ok {
			return
		}

		pcType, err := a.PcTypeService.PcType(r.Context(), req.TypeId)
		if err != nil {
			if errors.Is(err, mssql.ErrNotFound) {
				log.Warn("not found", sl.Err(err))
				response.NotFound(w)
			}
			log.Error("failed to get pc type", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, pcType)
	}
}
