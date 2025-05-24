package payment

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
}

func NewPaymentHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.POST("/payment/card", h.cardPayment)
}

func (h *Handler) cardPayment(ctx echo.Context) error {
	var req CardPayment
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	log.Info("Processing payment " + fmt.Sprint(req))

	return ctx.JSON(201, Status{Status: "Completed"})
}
