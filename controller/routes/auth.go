package routes

import (
	"falcon/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register
// @Summary Метод регистрации пользователя
// @Description Регистрация пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param RegisterInput body model.RegisterRequest true "Login data"
// @OperationId login
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/user [post]
func (w *WebApi) register(c *gin.Context) {
	var (
		ctx      = c.Request.Context()
		request  = model.RegisterRequest{}
		response = model.Response{}
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		w.logger.Errorf("received indalid data: %v with error: %v", request, err)
		response.ErrorResponse.Err = err
		response.ErrorResponse.Description = "please send valid data"
		response.ErrorResponse.StatusCode = model.BadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	registerResponse, appError := w.service.Register(ctx, request)
	if appError.Err != nil {
		w.logger.Error(appError.Message)
		if appError.StatusCode == model.BadRequest {
			response.ErrorResponse.Err = appError.Err
			response.ErrorResponse.Description = appError.Message
			response.ErrorResponse.StatusCode = appError.StatusCode
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response.ErrorResponse.Description = "something went wrong, please try later"
			response.ErrorResponse.Err = fmt.Errorf(response.ErrorResponse.Description)
			response.ErrorResponse.StatusCode = model.InternalServerError
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	response.Data = registerResponse

	c.JSON(http.StatusCreated, response)
}

// @Summary Метод входа пользователя
// @Description Вход пользователя под логином и паролем
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginInput body model.LoginRequest true "Login data"
// @OperationId login
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/login [post]
func (w *WebApi) login(c *gin.Context) {
	var (
		ctx      = c.Request.Context()
		request  = model.LoginRequest{}
		response = model.Response{}
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		w.logger.Errorf("received indalid data: %v with error: %v", request, err)
		response.ErrorResponse.Err = err
		response.ErrorResponse.Description = "please send valid data"
		response.ErrorResponse.StatusCode = model.BadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse, appError := w.service.Login(ctx, request)
	if appError.Err != nil {
		w.logger.Errorf(appError.Message)
		if appError.StatusCode == model.BadRequest {
			response.ErrorResponse.Err = appError.Err
			response.ErrorResponse.Description = appError.Message
			response.ErrorResponse.StatusCode = appError.StatusCode
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response.ErrorResponse.Description = "something went wrong, please try later"
			response.ErrorResponse.Err = fmt.Errorf(response.ErrorResponse.Description)
			response.ErrorResponse.StatusCode = model.InternalServerError
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	response.Data = loginResponse

	c.JSON(http.StatusOK, response)

}

// @Summary Метод генарации OTP
// @Description Генерация OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param GenerateOtp body model.GenerateOtpRequest true "Generate OTP data"
// @Success 200 {string} binary "PNG image data"
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/generate-otp [post]
func (w *WebApi) generateOtp(c *gin.Context) {
	var (
		ctx      = c.Request.Context()
		request  = &model.GenerateOtpRequest{}
		response = &model.Response{}
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		w.logger.Errorf("received indalid data: %v with error: %v", request, err)
		response.ErrorResponse.Err = err
		response.ErrorResponse.Description = "please send valid data"
		response.ErrorResponse.StatusCode = model.BadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	generateOtpResponse, appError := w.service.GenerateOTP(ctx, *request)
	if appError.Err != nil {
		w.logger.Errorf(appError.Message)
		if appError.StatusCode == model.BadRequest {
			response.ErrorResponse.Err = appError.Err
			response.ErrorResponse.Description = appError.Message
			response.ErrorResponse.StatusCode = appError.StatusCode
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response.ErrorResponse.Description = "something went wrong, please try later"
			response.ErrorResponse.Err = fmt.Errorf(response.ErrorResponse.Description)
			response.ErrorResponse.StatusCode = model.InternalServerError
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	c.Set("Content-Type", "image/png")
	c.Set("Content-Length", strconv.Itoa(len(generateOtpResponse.QrCode.Bytes())))

	c.Writer.Write(generateOtpResponse.QrCode.Bytes())

	c.Status(http.StatusOK)
}

// @Summary Метод валидации OTP
// @Description Валидация OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param ValidateOtp body model.ValidateOtpRequest true "Validate OTP data"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/validate-otp [post]
func (w *WebApi) validateOtp(c *gin.Context) {
	var (
		ctx      = c.Request.Context()
		request  = &model.ValidateOtpRequest{}
		response = &model.Response{}
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		w.logger.Errorf("received indalid data: %v with error: %v", request, err)
		response.ErrorResponse.Err = err
		response.ErrorResponse.Description = "please send valid data"
		response.ErrorResponse.StatusCode = model.BadRequest
		c.JSON(http.StatusBadRequest, response)
		return
	}

	validateOtpResponse, appError := w.service.ValidateOTP(ctx, *request)
	if appError.Err != nil {
		w.logger.Errorf(appError.Message)
		if appError.StatusCode == model.BadRequest {
			response.ErrorResponse.Err = appError.Err
			response.ErrorResponse.Description = appError.Message
			response.ErrorResponse.StatusCode = appError.StatusCode
			c.JSON(http.StatusBadRequest, response)
			return
		} else {
			response.ErrorResponse.Description = "something went wrong, please try later"
			response.ErrorResponse.Err = fmt.Errorf(response.ErrorResponse.Description)
			response.ErrorResponse.StatusCode = model.InternalServerError
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	response.Data = validateOtpResponse

	c.JSON(http.StatusOK, response)
}
