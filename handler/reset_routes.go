package handler

import "github.com/labstack/echo/v4"

func (h *Handler) fillResetRoutes(v1 *echo.Group) {
	v1.POST("reset", h.postReset)
}
