package use_cases

import (
	"context"
	"falconapi/domain/entities"
	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"strings"
)

func (uc *useCase) Register(ctx context.Context, request entities.RegisterRequest) (*entities.RegisterResponse, *entities.ErrorModel) {
	var (
		validate   = validator.New()
		errorModel entities.ErrorModel
		response   entities.RegisterResponse
	)

	err := validate.Struct(request)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 1
		return nil, &errorModel
	}

	var user = gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}
	if strings.TrimSpace(request.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{request.MobileNumber}
	}

	userResponse, err := uc.identityManager.CreateUser(ctx, user, request.Password, "viewer")
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 2
		return nil, &errorModel
	}

	response.User = userResponse

	return &response, &errorModel
}

func (uc *useCase) Login(ctx context.Context, request entities.LoginRequest) (*entities.LoginResponse, *entities.ErrorModel) {
	var (
		validate   = validator.New()
		response   entities.LoginResponse
		errorModel entities.ErrorModel
		attributes map[string][]string
	)

	err := validate.Struct(request)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 1
		return nil, &errorModel
	}

	user, _, err := uc.identityManager.LoginUser(ctx, request.Username, request.Password)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 1
		return nil, &errorModel
	}

	response.UserID = *user.ID

	if user.Attributes == nil {
		response.Message = "please generate OTP"
		response.OtpGenerated = false
	} else if user.Attributes != nil {
		attributes = *user.Attributes
		_, ok := attributes["otp_secret"]
		if !ok {
			response.Message = "please generate OTP"
			response.OtpGenerated = false
		} else if ok {
			response.Message = "please validate OTP"
			response.OtpGenerated = true
		}
	}

	return &response, &errorModel
}

func (uc *useCase) GenerateOTP(ctx context.Context, request entities.GenerateOtpRequest) (*entities.GenerateOtpResponse, *entities.ErrorModel) {
	var (
		errorModel entities.ErrorModel
		attribute  = make(map[string][]string)
		response   entities.GenerateOtpResponse
	)

	otpSecret, buf, err := uc.identityManager.GenerateOTP(ctx)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 2
		return nil, &errorModel
	}

	attribute["otp_secret"] = []string{otpSecret}

	err = uc.identityManager.UpdateUserAttribute(ctx, request.UserID, attribute)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 2
		return nil, &errorModel
	}

	response.UserID = request.UserID
	response.QrCode = buf

	return &response, &errorModel
}

func (uc *useCase) ValidateOTP(ctx context.Context, request entities.ValidateOtpRequest) (*entities.ValidateOtpResponse, *entities.ErrorModel) {
	var (
		errorModel entities.ErrorModel
		response   entities.ValidateOtpResponse
	)

	validOtp, err := uc.identityManager.ValidateOTP(ctx, request.UserID, request.OtpToken)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 2
		return nil, &errorModel
	}
	if !validOtp {
		errorModel.Message = "invalid otp"
		errorModel.Err = errors.New("invalid otp")
		errorModel.Code = 1
		return nil, &errorModel
	}

	user, token, err := uc.identityManager.LoginUser(ctx, request.Username, request.Password)
	if err != nil {
		errorModel.Message = err.Error()
		errorModel.Err = err
		errorModel.Code = 1
		return nil, &errorModel
	}

	response.Message = "valid OTP"
	response.UserID = *user.ID
	response.AccessToken = token.AccessToken
	response.RefreshToken = token.RefreshToken

	return &response, &errorModel
}
