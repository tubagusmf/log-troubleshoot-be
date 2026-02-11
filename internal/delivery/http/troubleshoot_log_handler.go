package http

import (
	"net/http"
	"strconv"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"

	"github.com/labstack/echo/v4"
)

type TroubleshootLogHandler struct {
	usecase model.ITroubleshootLogUsecase
}

func NewTroubleshootLogHandler(e *echo.Echo, troubleshootLogUsecase model.ITroubleshootLogUsecase) {
	handler := &TroubleshootLogHandler{
		usecase: troubleshootLogUsecase,
	}

	route := e.Group("v1/troubleshoot-log")
	route.POST("/create", handler.Create, AuthMiddleware)
	route.GET("/", handler.FindAll, AuthMiddleware)
	route.GET("/:id", handler.FindByID, AuthMiddleware)
	route.PUT("/update/:id", handler.Update, AuthMiddleware)
	route.DELETE("/delete/:id", handler.Delete, AuthMiddleware)
}

func (h *TroubleshootLogHandler) Create(c echo.Context) error {
	var input model.TroubleshootLog

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	data, err := h.usecase.Create(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *TroubleshootLogHandler) FindAll(c echo.Context) error {
	var filter model.TroubleshootLog

	filter.Status = c.QueryParam("status")
	filter.TicketNumber = c.QueryParam("ticket_number")

	data, err := h.usecase.FindAll(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *TroubleshootLogHandler) FindByID(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	data, err := h.usecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "data not found",
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (h *TroubleshootLogHandler) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var input model.TroubleshootLog
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if err := h.usecase.Update(c.Request().Context(), id, input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success update troubleshoot log",
	})
}

func (h *TroubleshootLogHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := h.usecase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success delete troubleshoot log",
	})
}
