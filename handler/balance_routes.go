package handler

import "github.com/labstack/echo/v4"

func (h *Handler) fillBalanceRoutes(v1 *echo.Group) {
	v1.GET("balance", h.getBalance)
}
