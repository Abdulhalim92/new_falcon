package model

import (
	"bytes"
)

type RegisterRequest struct {
	Username     string `json:"username" validate:"required,min=3,max=15"`
	Password     string `json:"password" validate:"required"`
	FirstName    string `json:"first_name" validate:"min=1,max=30"`
	LastName     string `json:"last_name" validate:"min=1,max=30"`
	Email        string `json:"email" validate:"required,email"`
	MobileNumber string `json:"mobile_number" validate:"required"`
	Role         string `json:"role" validate:"required"`
}

type RegisterResponse struct {
	UserID         int64
	KeycloakUserID string
	//User           *gocloak.User
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=15"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Message        string
	OtpGenerated   bool
	UserID         int64
	KeycloakUserID string
}

type ValidateOtpRequest struct {
	KeycloakUserID string `json:"keycloak_user_id"`
	OtpToken       string `json:"otp_token"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}

type ValidateOtpResponse struct {
	Message        string `json:"message,omitempty"`
	KeycloakUserID string `json:"keycloak_user_id,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	RefreshToken   string `json:"refresh_token,omitempty"`
}

type GenerateOtpRequest struct {
	KeycloakUserID string `json:"keycloak_user_id"`
}

type GenerateOtpResponse struct {
	KeycloakUserID string       `json:"keycloak_user_id,omitempty"`
	QrCode         bytes.Buffer `json:"qr_code,omitempty"`
}
