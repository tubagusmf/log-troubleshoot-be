package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type WorkTypeHandler struct {
	workTypeUsecase model.IWorkTypeUsecase
}

func NewWorkTypeHandler(e *echo.Echo, usecase model.IWorkTypeUsecase) {
	handler := &WorkTypeHandler{
		workTypeUsecase: usecase,
	}

	route := e.Group("v1/work-type")
	route.POST("/create", handler.Create, AuthMiddleware)
	route.GET("/", handler.FindAll, AuthMiddleware)
	route.GET("/:id", handler.FindByID, AuthMiddleware)
	route.PUT("/update/:id", handler.Update, AuthMiddleware)
	route.DELETE("/delete/:id", handler.Delete, AuthMiddleware)
}

func (h *WorkTypeHandler) Create(c echo.Context) error {
	var body model.CreateWorkTypeInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.workTypeUsecase.Create(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Work type created successfully",
		Data:    data,
	})
}

func (h *WorkTypeHandler) FindAll(c echo.Context) error {
	data, err := h.workTypeUsecase.FindAll(c.Request().Context(), model.WorkType{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   data,
	})
}

func (h *WorkTypeHandler) FindByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	data, err := h.workTypeUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   data,
	})
}

func (h *WorkTypeHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var body model.UpdateWorkTypeInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	if err := h.workTypeUsecase.Update(c.Request().Context(), id, body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Work type updated successfully",
		Data:    body,
	})
}

func (h *WorkTypeHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	if err := h.workTypeUsecase.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Work type deleted successfully",
	})
}
