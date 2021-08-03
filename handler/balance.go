package handler

import (
	"dt-tst/utils/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Get balance from an account_id
// @Summary Get balance from an account_id
// @Description Get balance from an account_id
// @Param account_id query int true "Account id"
// @Tags balance
// @Produce  json
// @Success 200 {number} float64
// @Failure 404 {number} int
// @Failure 500 {number} int
// @Router /balance [get]
func (h *Handler) getBalance(c echo.Context) (err error) {
	req := &types.BalanceRequest{}
	if err = c.Bind(req); err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("%d", 0))
	}
	if err = c.Validate(req); err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("%d", 0))
	}
	es, err := h.BalanceStore.GetBalance(req.AccountId)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("%d", 0))
	}
	return c.String(http.StatusOK, fmt.Sprintf("%.f", es))
}