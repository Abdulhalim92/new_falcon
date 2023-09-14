package routes

import (
	"falconapi/api"
	"falconapi/api/middlewares"
	"falconapi/use_cases"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type webApi struct {
	router  *gin.Engine
	useCase use_cases.UseCase
	token   string
}

func NewWebApi(engine *gin.Engine, useCase use_cases.UseCase, token string) api.WebApi {
	return &webApi{
		router:  engine,
		useCase: useCase,
		token:   token,
	}
}

func (w *webApi) InitRoutes() {

	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logrus.SetOutput(gin.DefaultWriter)

	w.router.Use(middlewares.CORSMiddleware())
	w.router.Use(gin.Recovery())
	w.router.Use(middlewares.Logger(logger))

	grp := w.router.Group("/v1")
	grp.GET("/", w.HealthCheck)
	grp.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := grp.Group("/auth")
	{
		auth.POST("/user", w.Register)
		auth.POST("/login", w.Login)
		auth.POST("/generate-otp", w.GenerateOtp)
		auth.POST("/validate-otp", w.ValidateOtp)
	}

	app := grp.Group("/api")
	app.Use(middlewares.JwtMiddleware())
	{
		app.GET("/", w.CheckMiddleware)
		app.GET("/terminals-statuses", w.GetTerminalsStatuses)
		app.GET("/terminals-info", w.GetTerminalsInfo)
		app.GET("/region", w.GetRegions)
	}
}
