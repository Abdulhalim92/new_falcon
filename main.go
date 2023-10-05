package main

import (
	"falcon/config"
	"falcon/contoller/routes"
	_ "falcon/docs"
	"falcon/domain/entity"
	gorm_db "falcon/pkg/gorm-db"
	"falcon/pkg/logging"
	"falcon/repository/database"
	"falcon/repository/identity"
	"falcon/service"
	"github.com/gin-gonic/gin"
)

// @title   Сервис админки FalconApi
// @version  1.0
// @description FalconApi

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host   localhost:8000
// @BasePath  /v1
// @schemes http https
func main() {

	config := config.ReadConfig()

	log := logging.GetLogger()

	router := gin.Default()

	db, err := gorm_db.InitDatabase(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})

	newIdentityManager := identity.NewIdentityManager(config.Keycloak, log)
	newUserRepo := database.NewUserRepo(log, db)

	newService := service.NewService(newIdentityManager, newUserRepo)

	newWebApi := routes.NewWebApi(router, log, newService)
	newWebApi.InitRoutes()

	//listenPort := config.AppParams.PortRun

	err = router.Run(":8000")
	log.Fatal(err)
}
