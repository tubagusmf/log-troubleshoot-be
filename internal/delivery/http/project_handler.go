package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
)

type ProjectHandler struct {
	projectUsecase model.IProjectUsecase
}

func NewProjectHandler(e *echo.Echo, projectUsecase model.IProjectUsecase) {
	handlers := &ProjectHandler{
		projectUsecase: projectUsecase,
	}

	routeProject := e.Group("v1/project")
	routeProject.POST("/create", handlers.Create, AuthMiddleware)
	routeProject.GET("/", handlers.FindAll, AuthMiddleware)
	routeProject.GET("/:id", handlers.FindByID, AuthMiddleware)
	routeProject.PUT("/update/:id", handlers.Update, AuthMiddleware)
	routeProject.DELETE("/delete/:id", handlers.Delete, AuthMiddleware)
}

func (handler *ProjectHandler) Create(c echo.Context) error {
	var body model.CreateProjectInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := handler.projectUsecase.Create(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Project created successfully",
		Data:    project,
	})
}

func (handler *ProjectHandler) FindAll(c echo.Context) error {
	projects, err := handler.projectUsecase.FindAll(c.Request().Context(), model.Project{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   projects,
	})
}

func (handler *ProjectHandler) FindByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	project, err := handler.projectUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status: http.StatusOK,
		Data:   project,
	})
}

func (handler *ProjectHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	var body model.UpdateProjectInput
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = handler.projectUsecase.Update(c.Request().Context(), id, body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Project updated successfully",
		Data:    body,
	})
}

func (handler *ProjectHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	claim, ok := c.Request().Context().Value(model.BearerAuthKey).(*model.CustomClaims)
	if !ok || claim == nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	err = handler.projectUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Project deleted successfully",
	})
}
