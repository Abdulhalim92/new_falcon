package routes

import (
	"falcon/domain/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (w *WebApi) register(c *gin.Context) {
	var (
		ctx      = c.Request.Context()
		request  = &model.RegisterRequest{}
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

	registerResponse, appError := w.service.Register(ctx, *request)
	if appError != nil {
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

	response.Data = registerResponse

	c.JSON(http.StatusCreated, response)
}

func (w *WebApi) login(c *gin.Context) {
	var (
		ctx      = c.Request.Context()
		request  = &model.LoginRequest{}
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

	loginResponse, appError := w.service.Login(ctx, *request)
	if appError != nil {
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
	if appError != nil {
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

	c.Status(http.StatusOK)
}

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
	if appError != nil {
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
