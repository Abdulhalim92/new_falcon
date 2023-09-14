package entities

import (
	"bytes"
	"github.com/Nerzal/gocloak/v13"
)

type RegisterRequest struct {
	Username     string `validate:"required,min=3,max=15"`
	Password     string `validate:"required"`
	FirstName    string `validate:"min=1,max=30"`
	LastName     string `validate:"min=1,max=30"`
	Email        string `validate:"required,email"`
	MobileNumber string
}

type RegisterResponse struct {
	User *gocloak.User
}

type LoginRequest struct {
	Username string `validate:"required,min=3,max=15"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	Message      string
	OtpGenerated bool
	UserID       string
}

type ValidateOtpRequest struct {
	UserID   string `json:"user_id"`
	OtpToken string `json:"otp_token"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ValidateOtpResponse struct {
	Message      string `json:"message,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type GenerateOtpRequest struct {
	UserID string `json:"user_id"`
}

type GenerateOtpResponse struct {
	UserID string       `json:"user_id,omitempty"`
	QrCode bytes.Buffer `json:"qr_code,omitempty"`
}
