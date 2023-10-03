package service

import (
	"bytes"
	"context"
	"falcon/domain/entity"
	"falcon/domain/model"
	"github.com/Nerzal/gocloak/v13"
)

type Identity interface {
	CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, *model.AppError)
	LoginUser(ctx context.Context, username, password string) (*gocloak.User, *gocloak.JWT, *model.AppError)
	UpdateUserAttribute(ctx context.Context, userID string, attribute map[string][]string) *model.AppError
	GenerateOTP(ctx context.Context) (string, bytes.Buffer, *model.AppError)
	ValidateOTP(ctx context.Context, userID, passcode string) (bool, *model.AppError)
	RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, *model.AppError)
}

type UserRepo interface {
	GetUserByID(ctx context.Context, userID int64) (*entity.User, *model.AppError)
	GetUserByName(ctx context.Context, username string) (*entity.User, *model.AppError)
	GetAllUsers(ctx context.Context) ([]entity.User, *model.AppError)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, *model.AppError)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, *model.AppError)
}

type Service interface {
	// Auth services
	Register(ctx context.Context, request model.RegisterRequest) (*model.RegisterResponse, *model.AppError)
	Login(ctx context.Context, request model.LoginRequest) (*model.LoginResponse, *model.AppError)
	GenerateOTP(ctx context.Context, request model.GenerateOtpRequest) (*model.GenerateOtpResponse, *model.AppError)
	ValidateOTP(ctx context.Context, request model.ValidateOtpRequest) (*model.ValidateOtpResponse, *model.AppError)
}
