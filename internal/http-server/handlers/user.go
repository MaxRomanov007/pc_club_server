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
	"pc_club_server/internal/services/auth"
	"pc_club_server/internal/services/user"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=32,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
type RegisterResponse struct {
	Access string `json:"access"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=32,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
type LoginResponse struct {
	Access string `json:"access_token"`
}

type RefreshResponse struct {
	Access string `json:"access_token"`
}

type UserResponse struct {
	Email   string
	Balance float32
}

func (a *API) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.SaveUser"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateRequest[RegisterRequest](w, r, log)
		if !ok {
			return
		}

		id, err := a.UserService.SaveUser(r.Context(), req.Email, req.Password)
		if err != nil {
			if errors.Is(err, user.ErrAlreadyExists) {
				log.Warn("user already exists", sl.Err(err))
				response.Conflict(w, "user is already existed")
				return
			}
			log.Error("failed to register user", sl.Err(err))
			response.Internal(w)
			return
		}

		access, refresh, err := a.AuthService.Tokens(r.Context(), id)
		if err != nil {
			var authError *auth.Error
			if ok := errors.As(err, &authError); ok {
				log.Warn("auth failed", sl.Err(err))
				response.AuthorizationFailed(w, authError)
				return
			}

			log.Error("failed to get tokens", sl.Err(err))
			response.Internal(w)
			return
		}

		response.SetRefreshCookie(w, a.Cfg.Auth, refresh)

		render.JSON(w, r, RegisterResponse{
			Access: access,
		})
	}
}

func (a *API) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.SaveUser"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		req, ok := request.DecodeAndValidateRequest[LoginRequest](w, r, log)
		if !ok {
			return
		}

		id, err := a.UserService.Login(r.Context(), req.Email, req.Password)
		if err != nil {
			if errors.Is(err, user.ErrInvalidCredentials) {
				log.Warn("invalid credentials", sl.Err(err))
				response.Unauthorized(w, "invalid credentials")
				return
			}
			log.Error("failed to login user", sl.Err(err))
			response.Internal(w)
			return
		}

		access, refresh, err := a.AuthService.Tokens(r.Context(), id)
		if err != nil {
			var authError *auth.Error
			if ok := errors.As(err, &authError); ok {
				log.Warn("auth failed", sl.Err(err))
				response.AuthorizationFailed(w, authError)
				return
			}

			log.Error("failed to get tokens", sl.Err(err))
			response.Internal(w)
			return
		}

		response.SetRefreshCookie(w, a.Cfg.Auth, refresh)

		render.JSON(w, r, LoginResponse{
			Access: access,
		})
	}
}

func (a *API) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.Refresh"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		refresh := request.RefreshToken(w, r, log, a.Cfg.Auth.Refresh.CookieName)
		if refresh == "" {
			return
		}

		access, refresh, err := a.AuthService.Refresh(r.Context(), refresh)
		if err != nil {
			var authError *auth.Error
			if ok := errors.As(err, &authError); ok {
				log.Warn("auth failed", sl.Err(err))
				response.AuthorizationFailed(w, authError)
				return
			}
			log.Error("failed to access token", sl.Err(err))
			response.Internal(w)
			return
		}

		response.SetRefreshCookie(w, a.Cfg.Auth, refresh)

		render.JSON(w, r, LoginResponse{
			Access: access,
		})
	}
}

func (a *API) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.Refresh"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		uid := request.MustUID(r)

		userData, err := a.UserService.User(r.Context(), uid)
		if err != nil {
			if errors.Is(err, user.ErrNotFound) {
				log.Warn("user not found", sl.Err(err))
				response.NotFound(w)
				return
			}
			log.Error("failed to get user", sl.Err(err))
			response.Internal(w)
			return
		}

		render.JSON(w, r, UserResponse{
			Email:   userData.Email,
			Balance: userData.Balance,
		})
	}
}

func (a *API) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.Logout"

		log := a.Log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		access := request.AccessToken(w, r, log)
		if access == "" {
			return
		}

		refreshToken := request.RefreshToken(w, r, log, a.Cfg.Auth.Refresh.CookieName)
		if refreshToken == "" {
			return
		}

		_, err := a.AuthService.BanTokens(r.Context(), access, refreshToken)
		if err != nil {
			var authError *auth.Error
			if ok := errors.As(err, &authError); ok {
				log.Warn("auth failed", sl.Err(err))
				response.AuthorizationFailed(w, authError)
				return
			}
			log.Error("failed to access token", sl.Err(err))
			response.Internal(w)
			return
		}
	}
}
