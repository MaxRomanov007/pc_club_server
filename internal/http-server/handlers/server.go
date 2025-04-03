package pcCLub

import (
	"context"
	"log/slog"
	"pc_club_server/internal/config"
	"pc_club_server/internal/domain/models"
)

type AuthService interface {
	Access(
		ctx context.Context,
		accessToken string,
	) (uid int64, err error)

	Refresh(
		ctx context.Context,
		refreshToken string,
	) (
		access string,
		refresh string,
		err error,
	)

	Tokens(
		ctx context.Context,
		uid int64,
	) (
		accessToken string,
		refreshToken string,
		err error,
	)

	BanTokens(
		ctx context.Context,
		accessToken string,
		refreshToken string,
	) (uid int64, err error)
}

type UserService interface {
	SaveUser(
		ctx context.Context,
		email string,
		password string,
	) (uid int64, err error)

	Login(
		ctx context.Context,
		email string,
		password string,
	) (uid int64, err error)

	User(
		ctx context.Context,
		uid int64,
	) (user models.User, err error)

	UserByEmail(
		ctx context.Context,
		email string,
	) (user models.User, err error)

	DeleteUser(
		ctx context.Context,
		uid int64,
	) (err error)

	IsAdmin(
		ctx context.Context,
		uid int64,
	) (err error)
}

type PcRoomService interface {
	PcRooms(
		ctx context.Context,
		pcTypeID int64,
	) (rooms []models.PcRoom, err error)
}

type PcTypeService interface {
	PcTypes(
		ctx context.Context,
		limit int,
		offset int,
	) (pcs []models.PcType, err error)

	PcType(
		ctx context.Context,
		typeID int64,
	) (pc models.PcType, err error)
}

type API struct {
	Log           *slog.Logger
	Cfg           *config.Config
	UserService   UserService
	AuthService   AuthService
	PcTypeService PcTypeService
	PcRoomService PcRoomService
}

func New(
	log *slog.Logger,
	cfg *config.Config,
	userService UserService,
	authService AuthService,
	pcTypeService PcTypeService,
	pcRoomService PcRoomService,
) *API {
	return &API{
		Log:           log,
		Cfg:           cfg,
		UserService:   userService,
		AuthService:   authService,
		PcTypeService: pcTypeService,
		PcRoomService: pcRoomService,
	}
}
