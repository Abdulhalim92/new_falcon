package routes

import (
	"falcon/contoller/middleware"
	"falcon/pkg/logging"
	"falcon/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type WebApi struct {
	logger  *logging.Logger
	router  *gin.Engine
	service service.Service
}

func NewWebApi(router *gin.Engine, logger *logging.Logger, service service.Service) *WebApi {
	return &WebApi{
		logger:  logger,
		router:  router,
		service: service,
	}
}

func (w *WebApi) InitRoutes() {

	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logrus.SetOutput(gin.DefaultWriter)

	w.router.Use(middleware.CORSMiddleware())
	w.router.Use(gin.Recovery())
	w.router.Use(middleware.Logger(logger))

	grp := w.router.Group("/v1")
	grp.GET("/", w.healthCheck)

	auth := grp.Group("/auth")
	{
		auth.POST("/user", w.register)
		auth.POST("/login", w.login)
		auth.POST("/generate-otp", w.generateOtp)
		auth.POST("/validate-otp", w.validateOtp)
	}
}

func (w *WebApi) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Состояние сервиса FalconApi"})
}
