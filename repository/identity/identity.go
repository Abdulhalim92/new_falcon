package identity

import (
	"bytes"
	"context"
	"falcon/config"
	"falcon/domain/model"
	"falcon/pkg/logging"
	"falcon/service"
	"fmt"
	"image/png"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pquerna/otp/totp"
)

type identityManager struct {
	baseUrl      string
	realm        string
	clientId     string
	clientSecret string
	logger       *logging.Logger
}

func NewIdentityManager(cfg config.KeycloakParams, logger *logging.Logger) service.Identity {
	return &identityManager{
		baseUrl:      cfg.BaseUrl,
		realm:        cfg.Realm,
		clientId:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		logger:       logger,
	}
}

func (im *identityManager) loginRestApiClient(ctx context.Context) (*gocloak.JWT, model.AppError) {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	token, err := client.LoginClient(ctx, im.clientId, im.clientSecret, im.realm)
	if err != nil {
		im.logger.Errorf("failed to login the client: %v", err)
		appError.Message = fmt.Sprintf("failed to login the client: %v", err)
		appError.StatusCode = model.Repository
		return nil, appError
	}
	return token, appError
}

func (im *identityManager) loginRestApiUser(ctx context.Context, username, password string) (*gocloak.JWT, model.AppError) {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	token, err := client.Login(ctx, im.clientId, im.clientSecret, im.realm, username, password)
	if err != nil {
		im.logger.Errorf("failed to login the user: %v", err)
		appError.Message = fmt.Sprintf("failed to login the user: %v", err)
		appError.StatusCode = model.Repository
		return nil, appError
	}

	return token, appError
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, model.AppError) {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	token, appError := im.loginRestApiClient(ctx)
	if appError.Err != nil {
		return nil, appError
	}

	userId, err := client.CreateUser(ctx, token.AccessToken, im.realm, user)
	if err != nil {
		im.logger.Errorf("failed to create the user: %v", err)
		appError.Err = err
		appError.Message = fmt.Sprintf("failed to create the user: %v", err)
		appError.StatusCode = model.BadRequest
		return nil, appError
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, im.realm, password, false)
	if err != nil {
		im.logger.Errorf("failed to set the password for the user: %v", err)
		appError.Message = fmt.Sprintf("failed to set the password for the user: %v", err)
		appError.StatusCode = model.Repository
		return nil, appError
	}

	var roleNameLowerCase = strings.ToLower(role)
	roleKeycloak, err := client.GetRealmRole(ctx, token.AccessToken, im.realm, roleNameLowerCase)
	if err != nil {
		im.logger.Errorf("failed to get role by name: '%v' with error: %v", roleNameLowerCase, err)
		appError.Message = fmt.Sprintf("failed to get role by name: '%v' with error: %v", roleNameLowerCase, err)
		appError.StatusCode = model.Repository
		return nil, appError
	}
	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.realm, userId, []gocloak.Role{
		*roleKeycloak,
	})
	if err != nil {
		im.logger.Errorf("failed to add a realm role to user: %v", err)
		appError.Message = fmt.Sprintf("failed to add a realm role to user: %v", err)
		appError.StatusCode = model.Repository
		return nil, appError
	}

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userId)
	if err != nil {
		im.logger.Errorf("failed to get recently created user: %v", err)
		appError.Message = fmt.Sprintf("failed to get recently created user: %v", err)
		appError.StatusCode = model.Repository
		return nil, appError
	}

	return userKeycloak, appError
}

func (im *identityManager) LoginUser(ctx context.Context, username, password string) (*gocloak.User, *gocloak.JWT, model.AppError) {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	userToken, appError := im.loginRestApiUser(ctx, username, password)
	if appError.Err != nil {
		return nil, nil, appError
	}

	clientToken, appError := im.loginRestApiClient(ctx)
	if appError.Err != nil {
		return nil, nil, appError
	}

	userInfo, err := client.GetUserInfo(ctx, userToken.AccessToken, im.realm)
	if err != nil {
		im.logger.Errorf("failed to get user information: %v", err)
		appError.Message = fmt.Sprintf("failed to get user information: %v", err)
		appError.StatusCode = model.Repository
		return nil, nil, appError
	}

	user, err := client.GetUserByID(ctx, clientToken.AccessToken, im.realm, *userInfo.Sub)
	if err != nil {
		im.logger.Errorf("failed to get user by ID: %v", err)
		appError.Message = fmt.Sprintf("failed to get user by ID: %v", err)
		appError.StatusCode = model.Repository
		return nil, nil, appError
	}

	return user, userToken, appError
}

func (im *identityManager) UpdateUserAttribute(ctx context.Context, userID string, attribute map[string][]string) model.AppError {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	clientToken, appError := im.loginRestApiClient(ctx)
	if appError.Err != nil {
		return appError
	}

	user, err := client.GetUserByID(ctx, clientToken.AccessToken, im.realm, userID)
	if err != nil {
		im.logger.Errorf("failed to ge user: %v", err)
		appError.Message = fmt.Sprintf("failed to ge user: %v", err)
		appError.StatusCode = model.Repository
		return appError
	}

	user.Attributes = &attribute

	err = client.UpdateUser(ctx, clientToken.AccessToken, im.realm, *user)
	if err != nil {
		im.logger.Errorf("failed to update user: %v", err)
		appError.Message = fmt.Sprintf("failed to update user: %v", err)
		appError.StatusCode = model.Repository
		return appError
	}

	return appError
}

func (im *identityManager) GenerateOTP(ctx context.Context) (string, bytes.Buffer, model.AppError) {
	var (
		qrCode   bytes.Buffer
		appError model.AppError
	)

	// генерация OTP
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "humopay",
		AccountName: "humopay",
		SecretSize:  15,
	})
	if err != nil {
		im.logger.Errorf("failed to generate OTP: %v", err)
		appError.Message = fmt.Sprintf("failed to generate OTP: %v", err)
		appError.StatusCode = model.Repository
		return "", bytes.Buffer{}, appError
	}

	// создание QR - кода
	img, err := key.Image(200, 200)
	if err != nil {
		im.logger.Errorf("failed to create image for QR - code: %v", err)
		appError.Message = fmt.Sprintf("failed to create image for QR - code: %v", err)
		appError.StatusCode = model.Repository
		return "", bytes.Buffer{}, appError
	}

	err = png.Encode(&qrCode, img)
	if err != nil {
		im.logger.Errorf("failed to save image from created QR - code: %v", err)
		appError.Message = fmt.Sprintf("failed to save image from created QR - code: %v", err)
		appError.StatusCode = model.Repository
		return "", bytes.Buffer{}, appError
	}

	return key.Secret(), qrCode, appError
}

func (im *identityManager) ValidateOTP(ctx context.Context, userID, passcode string) (bool, model.AppError) {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	clientToken, appError := im.loginRestApiClient(ctx)
	if appError.Err != nil {
		return false, appError
	}

	user, err := client.GetUserByID(ctx, clientToken.AccessToken, im.realm, userID)
	if err != nil {
		im.logger.Errorf("failed to get user by ID: %v", err)
		appError.Message = fmt.Sprintf("failed to get user by ID: %v", err)
		appError.StatusCode = model.Repository
		return false, appError
	}

	var otpSecret string

	if user.Attributes == nil {
		im.logger.Errorf("for this user didn't generated OTP")
		appError.Message = fmt.Sprintf("for this user didn't generated OTP")
		appError.StatusCode = model.Repository
		return false, appError
	} else if user.Attributes != nil {
		otpSecret = (*user.Attributes)["otp_secret"][0]
		if otpSecret == "" {
			im.logger.Errorf("for this user didn't generated OTP")
			appError.Message = fmt.Sprintf("for this user didn't generated OTP")
			appError.StatusCode = model.Repository
			return false, appError
		}
	}

	otpSecret = (*user.Attributes)["otp_secret"][0]

	validate := totp.Validate(passcode, otpSecret)

	return validate, appError
}

func (im *identityManager) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, model.AppError) {
	var (
		client   = gocloak.NewClient(im.baseUrl)
		appError model.AppError
	)

	rptResult, err := client.RetrospectToken(ctx, accessToken, im.clientId, im.clientSecret, im.realm)
	if err != nil {
		im.logger.Errorf("failed to retrospect token: %v", err)
		appError.Message = fmt.Sprintf("failed to retrospect token: %v", err)
		appError.StatusCode = model.Repository
		return nil, appError
	}
	return rptResult, appError
}
