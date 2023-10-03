package service

import (
	"context"
	"falcon/domain/entity"
	"falcon/domain/model"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator/v10"
)

func (s *service) Register(ctx context.Context, request model.RegisterRequest) (*model.RegisterResponse, *model.AppError) {
	var (
		validate = validator.New()
		appError *model.AppError
		response *model.RegisterResponse
		user     *entity.User
	)

	err := validate.Struct(response)
	if err != nil {
		appError.Err = err
		appError.Message = err.Error()
		appError.StatusCode = model.BadRequest
		return nil, appError
	}

	// adding user to keycloak
	keycloakUser := &gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(false),
		Attributes:    &map[string][]string{},
	}
	(*keycloakUser.Attributes)["mobile"] = []string{request.MobileNumber}

	keycloakUser, appError = s.identity.CreateUser(ctx, *keycloakUser, request.Password, request.Role)
	if appError != nil {
		return nil, appError
	}

	// check user for existing in database
	user, appError = s.userRepo.GetUserByName(ctx, request.Username)
	if appError != nil {
		return nil, appError
	}
	if user.ID != 0 {
		appError.Message = fmt.Sprintf("user with name %v is already exists", request.Username)
		appError.StatusCode = model.BadRequest
		return nil, appError
	}

	// adding user to database
	user.KeycloakID = *keycloakUser.ID
	user.Username = request.Username
	user.Role = request.Role
	user.Email = request.Email
	user.FullName = request.FirstName + " " + request.LastName

	user, appError = s.userRepo.CreateUser(ctx, user)
	if appError != nil {
		return nil, appError
	}

	return &model.RegisterResponse{
		UserID:         user.ID,
		KeycloakUserID: user.KeycloakID,
	}, nil
}

func (s *service) Login(ctx context.Context, request model.LoginRequest) (*model.LoginResponse, *model.AppError) {
	var (
		validate   = validator.New()
		response   *model.LoginResponse
		appError   *model.AppError
		attributes map[string][]string
	)

	err := validate.Struct(response)
	if err != nil {
		appError.Message = err.Error()
		appError.Err = err
		appError.StatusCode = model.BadRequest
		return nil, appError
	}

	keycloakUser, _, appError := s.identity.LoginUser(ctx, request.Username, request.Password)
	if appError != nil {
		return nil, appError
	}

	response.KeycloakUserID = *keycloakUser.ID

	if keycloakUser.Attributes == nil {
		response.Message = "please generate OTP"
		response.OtpGenerated = false
	} else if keycloakUser.Attributes != nil {
		attributes = *keycloakUser.Attributes
		_, ok := attributes["otp_secret"]
		if !ok {
			response.Message = "please generate OTP"
			response.OtpGenerated = false
		} else if ok {
			response.Message = "please validate OTP"
			response.OtpGenerated = true
		}
	}

	return response, nil
}

func (s *service) GenerateOTP(ctx context.Context, request model.GenerateOtpRequest) (*model.GenerateOtpResponse, *model.AppError) {
	var (
		appError  *model.AppError
		response  *model.GenerateOtpResponse
		attribute = make(map[string][]string)
	)

	otpSecret, buf, appError := s.identity.GenerateOTP(ctx)
	if appError != nil {
		return nil, appError
	}

	attribute["otp_secret"] = []string{otpSecret}

	appError = s.identity.UpdateUserAttribute(ctx, request.KeycloakUserID, attribute)
	if appError != nil {
		return nil, appError
	}

	response.KeycloakUserID = request.KeycloakUserID
	response.QrCode = buf

	return response, nil
}

func (s *service) ValidateOTP(ctx context.Context, request model.ValidateOtpRequest) (*model.ValidateOtpResponse, *model.AppError) {
	var (
		appError *model.AppError
		response *model.ValidateOtpResponse
	)

	validOTP, appError := s.identity.ValidateOTP(ctx, request.Username, request.Password)
	if appError != nil {
		return nil, appError
	}
	if !validOTP {
		appError.Message = "invalid otp"
		appError.Err = fmt.Errorf(appError.Message)
		appError.StatusCode = model.BadRequest
		return nil, appError
	}

	keycloakUser, token, appError := s.identity.LoginUser(ctx, request.Username, request.Password)
	if appError != nil {
		return nil, appError
	}

	response.Message = "valid OTP"
	response.KeycloakUserID = *keycloakUser.ID
	response.AccessToken = token.AccessToken
	response.RefreshToken = token.RefreshToken

	return response, nil
}
