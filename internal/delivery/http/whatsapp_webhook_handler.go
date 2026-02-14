package http

import (
	"fmt"
	"net/http"

	"github.com/tubagusmf/log-troubleshoot-be/internal/model"
	"github.com/tubagusmf/log-troubleshoot-be/internal/usecase"

	"github.com/labstack/echo/v4"
)

type WhatsAppWebhookHandler struct {
	usecase *usecase.WhatsAppConsumerUsecase
}

func NewWhatsAppWebhookHandler(
	e *echo.Echo,
	usecase *usecase.WhatsAppConsumerUsecase,
) {
	handler := &WhatsAppWebhookHandler{usecase: usecase}

	e.POST("webhook/whatsapp", handler.Handle)
}

func (h *WhatsAppWebhookHandler) Handle(c echo.Context) error {
	var payload model.WhatsAppWebhookRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	err := h.usecase.Consume(c.Request().Context(), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "webhook received",
	})
}

func (h *WhatsAppWebhookHandler) ReceiveWebhook(c echo.Context) error {
	var payload map[string]interface{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	}

	fmt.Println("Incoming Payload:", payload)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "webhook received",
	})
}
