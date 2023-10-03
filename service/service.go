package service

type service struct {
	identity Identity
	userRepo UserRepo
}

func NewService(identity Identity, userRepo UserRepo) Service {
	return &service{
		identity: identity,
		userRepo: userRepo,
	}
}
