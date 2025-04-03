package pcClub

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log/slog"
	"net/http"
	"pc_club_server/internal/config"
	"pc_club_server/internal/http-server/handlers"
	"pc_club_server/internal/http-server/middleware/auth/authAdmin"
	"pc_club_server/internal/http-server/middleware/auth/authorization"
	"pc_club_server/internal/http-server/middleware/logger"
	"pc_club_server/internal/lib/api/logger/sl"
)

type App struct {
	Log         *slog.Logger
	HTTPSServer *http.Server
	Cfg         *config.HTTPSServerConfig
}

func New(cfg *config.HTTPSServerConfig, api *pcCLub.API) *App {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Depth", "UserService-Agent", "X-File-Size", "X-Requested-With", "If-Modified-Since", "X-File-Name", "Cache-Control", "Access-Control-Expose-Headers", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(logger.New(api.Log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Post("/register", api.Register())
	r.Post("/login", api.Login())
	r.Post("/refresh", api.Refresh())
	r.Post("/logout", api.Logout())

	r.Get("/pc-types", api.PcTypes())
	r.Get("/pc-types/{type-id}", api.PcType())

	r.Get("/pc-rooms", api.PcRooms())

	//routes to be authorized
	r.Group(func(r chi.Router) {
		r.Use(authorization.Authorize(api.Log, api.AuthService))

		r.Post("/user", api.User())
	})

	//admin routes
	r.Group(func(r chi.Router) {
		r.Use(authAdmin.AuthAdmin(api.Log, api.AuthService, api.UserService))

	})

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &App{
		Log:         api.Log,
		HTTPSServer: srv,
		Cfg:         cfg,
	}
}

func (a *App) MustRun() {
	const op = "app.club.MustRun"

	log := a.Log.With(
		slog.String("operation", op),
	)

	if err := a.RunClub(); err != nil {
		log.Error("failed to start server", sl.Err(err))

		panic(err)
	}
}

func (a *App) RunClub() error {
	const op = "app.club.Run"

	if err := a.HTTPSServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: failed to start club server: %w", op, err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	const op = "app.cars.Run"

	err := a.HTTPSServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("%s: failed to stop club server: %w", op, err)
	}

	return nil
}
