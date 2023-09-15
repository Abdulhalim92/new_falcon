package database

import (
	"falconapi/pkg/logging"
	"falconapi/use_cases"
	"gorm.io/gorm"
)

type database struct {
	log    *logging.Logger
	dbConn *gorm.DB
}

func NewDatabase(log *logging.Logger, db *gorm.DB) use_cases.TerminalRepository {
	return &database{
		log:    log,
		dbConn: db,
	}
}
