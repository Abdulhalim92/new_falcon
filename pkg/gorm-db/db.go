package gorm_db

import (
	"falcon/config"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase(cfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Dushanbe",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Dbname, cfg.SslMode)

	log.Println(dsn)

	dbConn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Connection success-->> host:port -->> %s:%s", viper.GetString("Database.Host"), viper.GetString("Database.Port"))

	return dbConn, nil
}
