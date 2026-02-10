package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type LocationHandler struct {
	locationUsecase model.ILocationUsecase
}

func NewLocationHandler(e *echo.Echo, locationUsecase model.ILocationUsecase) {
	handler := &LocationHandler{
		locationUsecase: locationUsecase,
	}

	route := e.Group("v1/location")
	route.POST("/create", handler.Create, AuthMiddleware)
	route.GET("/", handler.FindAll, AuthMiddleware)
	route.GET("/:id", handler.FindByID, AuthMiddleware)
	route.PUT("/update/:id", handler.Update, AuthMiddleware)
	route.DELETE("/delete/:id", handler.Delete, AuthMiddleware)
}

func (h *LocationHandler) Create(c echo.Context) error {
	var body model.CreateLocationInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	location, err := h.locationUsecase.Create(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Location created successfully",
		Data:    location,
	})
}

func (h *LocationHandler) FindAll(c echo.Context) error {
	locations, err := h.locationUsecase.FindAll(c.Request().Context(), model.Location{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   locations,
	})
}

func (h *LocationHandler) FindByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	location, err := h.locationUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   location,
	})
}

func (h *LocationHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	var body model.UpdateLocationInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = h.locationUsecase.Update(c.Request().Context(), id, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Location updated successfully",
		Data:    body,
	})
}

func (h *LocationHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = h.locationUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Location deleted successfully",
	})
}
