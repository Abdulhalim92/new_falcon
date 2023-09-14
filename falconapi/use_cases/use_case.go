package use_cases

type useCase struct {
	identityManager    IdentityManager
	terminalRepository TerminalRepository
}

func NewUseCase(manager IdentityManager, terminalRepository TerminalRepository) UseCase {
	return &useCase{
		identityManager:    manager,
		terminalRepository: terminalRepository,
	}
}
