package routes

import (
	"falcon/controller/middleware"
	"falcon/pkg/logging"
	"falcon/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	grp.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := grp.Group("/auth")
	{
		auth.POST("/user", w.register)
		auth.POST("/login", w.login)
		auth.POST("/generate-otp", w.generateOtp)
		auth.POST("/validate-otp", w.validateOtp)
	}
}

// @Summary Метод проверки состояния системы
// @Description Проверка состояния системы
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router / [get]
func (w *WebApi) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Состояние сервиса FalconApi"})
}
