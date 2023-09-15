package routes

import (
	"falconapi/domain/entities"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary Проверка работоспособности системы
// @Description Проверка работоспособности системы
// @Tags HealthCheck
// @Accept json
// @Produce json
// @OperationId healthCheck
// @Success 200 {object} string
// @Router / [get]
func (w *webApi) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Состояние сервиса FalconApi"})
}

// @Summary Метод регистрации пользователя
// @Description Регистрация пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param RegisterInput body entities.RegisterRequest true "Login data"
// @OperationId login
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/user [post]
func (w *webApi) Register(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		requset = entities.RegisterRequest{}
	)

	err := c.BindJSON(&requset)
	if err != nil {
		log.Println(err, "unable to parse incoming request")
		c.JSON(http.StatusBadRequest, "unable to parse incoming request")
		return
	}

	response, errorModel := w.useCase.Register(ctx, requset)
	if errorModel != nil {
		log.Println(errorModel.Err, "unable to register user")
		if errorModel.Code == 1 {
			c.JSON(http.StatusBadRequest, "please send valid data")
			return
		} else if errorModel.Code == 2 {
			c.JSON(http.StatusInternalServerError, "something went wrong")
			return
		}
	}

	c.JSON(http.StatusCreated, response.User.ID)
}

// @Summary Метод входа пользователя
// @Description Вход пользователя под логином и паролем
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginInput body entities.LoginRequest true "Login data"
// @OperationId login
// @Success 200 {object} entities.LoginResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/login [post]
func (w *webApi) Login(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request = entities.LoginRequest{}
	)

	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err, "unable to parse incoming request")
		c.JSON(http.StatusBadRequest, "unable to parse incoming request")
		return
	}

	response, errorModel := w.useCase.Login(ctx, request)
	if errorModel != nil {
		log.Println(errorModel.Err, "unable to login user")
		if errorModel.Code == 1 {
			c.JSON(http.StatusBadRequest, "incorrect login or password")
			return
		} else if errorModel.Code == 2 {
			c.JSON(http.StatusInternalServerError, "something went wrong")
			return
		}
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Метод генарации OTP
// @Description Генерация OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param GenerateOtp body entities.GenerateOtpRequest true "Generate OTP data"
// @Success 200 {string} binary "PNG image data"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/generate-otp [post]
func (w *webApi) GenerateOtp(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request = entities.GenerateOtpRequest{}
	)

	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err, "unable to parse incoming request")
		c.JSON(http.StatusBadRequest, "unable to parse incoming request")
		return
	}

	response, errorModel := w.useCase.GenerateOTP(ctx, request)
	if errorModel != nil {
		log.Println(errorModel.Err, "unable to generate OTP")
		if errorModel.Code == 1 {
			c.JSON(http.StatusBadRequest, "please send valid data")
			return
		} else if errorModel.Code == 2 {
			c.JSON(http.StatusInternalServerError, "something went wrong")
			return
		}
	}

	c.Set("Content-Type", "image/png")
	c.Set("Content-Length", strconv.Itoa(len(response.QrCode.Bytes())))

	c.Status(http.StatusOK)

	_, err = c.Writer.Write(response.QrCode.Bytes())
	if err != nil {
		log.Println("unable to write image.")
	}
}

// @Summary Метод валидации OTP
// @Description Валидация OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param ValidateOtp body entities.ValidateOtpRequest true "Validate OTP data"
// @Success 200 {object} entities.ValidateOtpResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/validate-otp [post]
func (w *webApi) ValidateOtp(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		request = entities.ValidateOtpRequest{}
	)

	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err, "unable to parse incoming request")
		c.JSON(http.StatusBadRequest, "unable to parse incoming request")
		return
	}

	response, errorModel := w.useCase.ValidateOTP(ctx, request)
	if errorModel != nil {
		log.Println(err, "unable to validate OTP")
		if errorModel.Code == 1 {
			c.JSON(http.StatusBadRequest, "please send valid data")
			return
		} else if errorModel.Code == 2 {
			c.JSON(http.StatusInternalServerError, "something went wrong")
			return
		}
	}

	c.JSON(http.StatusOK, response)
}
