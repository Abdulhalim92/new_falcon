package use_cases

import (
	"bytes"
	"context"
	"falconapi/domain/entities"
	"github.com/Nerzal/gocloak/v13"
	"github.com/google/uuid"
)

type TerminalRepository interface {
	GetAll(ctx context.Context) ([]entities.TerminalStatus, error)
	GetRegions(ctx context.Context) ([]entities.TRegion, error)
	GetAllWithRegionsNames(ctx context.Context) ([]entities.TerminalStatus, error)
	GetTerminalByID(ctx context.Context, id string) (*entities.TEndpoint, error)
	GenerateTermPassword(ctx context.Context, userID string, endpoint entities.TEndpoint) (*entities.TEndpoint, error)
	CountTerminalByNum(ctx context.Context, incomingTerminalData entities.TerminalCreate) (int64, error)
	NewTerminal(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error)
	UpdateTerminal(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error)
	ChangeApi(ctx context.Context, newApi *uuid.UUID, endpoint entities.TEndpoint) (*entities.TEndpoint, error)
	GetXMLGate(ctx context.Context) ([]entities.TEndpoint, error)
	CreateXMLGate(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error)
	UpdateTerminalXML(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error)
	GetStatusPingByID(ctx context.Context, id string) (*entities.StatusTerminal, error)
	GetAllCommands(ctx context.Context) ([]entities.EndpointCommand, error)
	CreateCommand(ctx context.Context, cmd entities.EndpointCommand) (*entities.EndpointCommand, error)
	GetCommandsByDate(ctx context.Context, cmdSearch entities.CommandSearch) ([]entities.EndpointCommand, error)
	TerminalReportKiosk(ctx context.Context, report entities.IncomingTermReport) ([]entities.TermReportKiosk, error)
	TerminalReport(ctx context.Context, report entities.IncomingTermReport) ([]entities.TermReport, error)
	GetTerminalsByAgent(ctx context.Context, agentID int64) ([]entities.TEndpoint, error)
	CreateCommandAuto(ctx context.Context) error
}

type IdentityManager interface {
	CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error)
	LoginUser(ctx context.Context, username, password string) (*gocloak.User, *gocloak.JWT, error)
	UpdateUserAttribute(ctx context.Context, userID string, attribute map[string][]string) error
	GenerateOTP(ctx context.Context) (string, bytes.Buffer, error)
	ValidateOTP(ctx context.Context, userID, passcode string) (bool, error)
}

type TokenRetrospect interface {
	RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
}

type UseCase interface {
	// Auth services
	Register(ctx context.Context, request entities.RegisterRequest) (*entities.RegisterResponse, *entities.ErrorModel)
	Login(ctx context.Context, request entities.LoginRequest) (*entities.LoginResponse, *entities.ErrorModel)
	GenerateOTP(ctx context.Context, request entities.GenerateOtpRequest) (*entities.GenerateOtpResponse, *entities.ErrorModel)
	ValidateOTP(ctx context.Context, request entities.ValidateOtpRequest) (*entities.ValidateOtpResponse, *entities.ErrorModel)

	// Terminals services
	GetRegions(ctx context.Context) ([]entities.TRegion, error)
	GetTerminalsInfo(ctx context.Context) ([]entities.TerminalStatus, error)
	GetTerminalsStatuses(ctx context.Context) ([]entities.TerminalStatus, error)
}
