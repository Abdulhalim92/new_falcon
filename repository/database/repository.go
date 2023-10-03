package database

import (
	"falcon/pkg/logging"
	"falcon/service"
	"gorm.io/gorm"
)

type repository struct {
	logger *logging.Logger
	db     *gorm.DB
}

func NewUserRepo(logger *logging.Logger, db *gorm.DB) service.UserRepo {
	return &repository{
		logger: logger,
		db:     db,
	}
}
