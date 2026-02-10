package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type DeviceHandler struct {
	deviceUsecase model.IDeviceUsecase
}

func NewDeviceHandler(e *echo.Echo, deviceUsecase model.IDeviceUsecase) {
	handler := &DeviceHandler{
		deviceUsecase: deviceUsecase,
	}

	route := e.Group("v1/device")
	route.POST("/create", handler.Create, AuthMiddleware)
	route.GET("/", handler.FindAll, AuthMiddleware)
	route.GET("/:id", handler.FindByID, AuthMiddleware)
	route.PUT("/update/:id", handler.Update, AuthMiddleware)
	route.DELETE("/delete/:id", handler.Delete, AuthMiddleware)
}

func (h *DeviceHandler) Create(c echo.Context) error {
	var body model.CreateDeviceInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	device, err := h.deviceUsecase.Create(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Device created successfully",
		Data:    device,
	})
}

func (h *DeviceHandler) FindAll(c echo.Context) error {
	devices, err := h.deviceUsecase.FindAll(c.Request().Context(), model.Device{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   devices,
	})
}

func (h *DeviceHandler) FindByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	device, err := h.deviceUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   device,
	})
}

func (h *DeviceHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	var body model.UpdateDeviceInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = h.deviceUsecase.Update(c.Request().Context(), id, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Device updated successfully",
		Data:    body,
	})
}

func (h *DeviceHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = h.deviceUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Device deleted successfully",
	})
}
