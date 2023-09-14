package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary Метод проверки middleware
// @Description Проверка middleware
// @Tags Terminals
// @OperationId checkMiddleware
// @Success 200 {string} 1
// @Router /api/ [get]
func (w *webApi) CheckMiddleware(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "token is working"})
}

// @Summary Метод получения статусов терминалов
// @Description Получение статусов терминалов
// @Tags Terminals
// @Accept json
// @Produce json
// @OperationId getTerminalsStatuses
// @Success 200 {object} map[string][]entities.TerminalStatus
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/terminalstatuses [get]
func (w *webApi) GetTerminalsStatuses(c *gin.Context) {
	ctx := c.Request.Context()

	terminalsStatuses, err := w.useCase.GetTerminalsStatuses(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	c.JSON(http.StatusOK, gin.H{"data": terminalsStatuses})
}

// @Summary Метод получения инфо - статусов терминалов
// @Description Получение инфо - статусов терминалов
// @Tags Terminals
// @Accept json
// @Produce json
// @OperationId getTerminalsInfo
// @Success 200 {object} map[string][]entities.TerminalStatus
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/terminalsinfo [get]
func (w *webApi) GetTerminalsInfo(c *gin.Context) {
	ctx := c.Request.Context()

	terminalsInfo, err := w.useCase.GetTerminalsInfo(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": terminalsInfo})
}

// @Summary Метод получения списка регионов
// @Description Получение списка регионов
// @Tags Terminals
// @Accept json
// @Produce json
// @OperationId getRegions
// @Success 200 {object} map[string][]entities.TRegion
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/region [get]
func (w *webApi) GetRegions(c *gin.Context) {
	ctx := c.Request.Context()

	regions, err := w.useCase.GetRegions(ctx)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": regions})
}
