package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	h.fillBalanceRoutes(v1)
	h.fillEventRoutes(v1)
	h.fillResetRoutes(v1)
}
