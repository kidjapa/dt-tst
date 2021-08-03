package handler

import "github.com/labstack/echo/v4"

func (h *Handler) fillEventRoutes(v1 *echo.Group) {
	v1.POST("event", h.postEvent)
}
