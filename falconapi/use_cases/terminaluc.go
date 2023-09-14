package use_cases

import (
	"context"
	"falconapi/domain/entities"
)

func (uc *useCase) GetRegions(ctx context.Context) ([]entities.TRegion, error) {
	return uc.terminalRepository.GetRegions(ctx)
}

func (uc *useCase) GetTerminalsInfo(ctx context.Context) ([]entities.TerminalStatus, error) {
	return uc.terminalRepository.GetAllWithRegionsNames(ctx)
}

func (uc *useCase) GetTerminalsStatuses(ctx context.Context) ([]entities.TerminalStatus, error) {
	return uc.terminalRepository.GetAll(ctx)
}
