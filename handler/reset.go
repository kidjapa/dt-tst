package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Reset state before starting tests
// @Summary Reset state before starting tests
// @Description Reset state before starting tests
// @Tags reset
// @Produce  json
// @Success 200 {string} string
// @Failure 404 {number} int
// @Router /reset [post]
func (h *Handler) postReset(c echo.Context) (err error) {
	res, _ := h.Reset.Reset()
	if !res {
		return c.String(http.StatusNotFound, "NOK")
	}
	return c.String(http.StatusOK, "OK")
}