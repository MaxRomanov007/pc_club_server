package response

import (
	"github.com/go-playground/validator"
	"net/http"
	"pc_club_server/internal/config"
	validator2 "pc_club_server/internal/lib/api/validator"
	"pc_club_server/internal/lib/cookie"
	"pc_club_server/internal/services/auth"
	"strings"
)

func Internal(w http.ResponseWriter) {
	http.Error(w, "internal error", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	http.Error(w, "not found", http.StatusUnauthorized)
}

func Conflict(w http.ResponseWriter, message ...string) {
	http.Error(w, strings.Join(message, " "), http.StatusConflict)
}

func Unauthorized(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusUnauthorized)
}

func ValidationFailed[T any](w http.ResponseWriter, errs validator.ValidationErrors) {
	http.Error(w, validator2.ValidationError[T](errs), http.StatusBadRequest)
}

func AuthorizationFailed(w http.ResponseWriter, err *auth.Error) {
	http.Error(w, err.Error(), http.StatusUnauthorized)
}

func SetRefreshCookie(w http.ResponseWriter, cfg *config.AuthConfig, refreshToken string) {
	cookie.Set(
		w,
		cfg.Refresh.CookieName,
		refreshToken,
		cfg.UrlPath,
		cfg.Refresh.TTL,
	)
}
