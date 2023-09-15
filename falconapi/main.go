package main

import (
	"falconapi/controller/routes"
	_ "falconapi/docs"
	gorm_db "falconapi/pkg/gorm-db"
	"falconapi/pkg/logging"
	"falconapi/repository/database"
	"falconapi/repository/identity"
	"falconapi/use_cases"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title   Сервис админки FalconApi
// @version  1.0
// @description FalconApi

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host   127.0.0.1:3006
// @BasePath  /v1
// @schemes http
func main() {

	log := logging.GetLogger()

	initViper()

	router := gin.Default()
	db, err := gorm_db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	newIdentityManager := identity.NewIdentityManager()
	newDatabase := database.NewDatabase(log, db)

	newUseCase := use_cases.NewUseCase(newIdentityManager, newDatabase)

	newWebApi := routes.NewWebApi(router, newUseCase, "")
	newWebApi.InitRoutes()

	var listenIp = viper.GetString("ListenIP")
	var listenPort = viper.GetString("ListenPort")

	log.Printf("will listen on %v:%v", listenIp, listenPort)

	err = router.Run(fmt.Sprintf("%v:%v", listenIp, listenPort))
	log.Fatal(err)
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("demo")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("viper config initialized")
}
