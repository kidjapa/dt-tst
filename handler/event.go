package handler

import (
	"dt-tst/utils"
	"dt-tst/utils/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

// PostEvent Post an event
// @Summary PostEvent Post an event
// @Description PostEvent Post an event
// @Param event body types.EventRequest true "spritesheet with width and height of the expected sprite cut"
// @Tags event
// @Produce  json
// @Success 201 {object} types.PostEventCommonResponse
// @Failure 404 {number} int
// @Router /event [post]
func (h *Handler) postEvent(c echo.Context) (err error) {

	req := &types.EventRequest{}
	if err = c.Bind(req); err != nil {
		return utils.GetDefaultErrorMessage(c)
	}
	if err = c.Validate(req); err != nil {
		return utils.GetDefaultErrorMessage(c)
	}

	switch req.Type {
	case types.DepositRequestTypeDeposit:
		res := &types.DepositResponse{}
		res, err = h.EventStore.PostDeposit(req)
		if err != nil {
			return utils.GetDefaultErrorMessage(c)
		}
		return c.JSON(http.StatusCreated, res)
	case types.DepositRequestTypeWithDraw:
		res := &types.WithdrawResponse{}
		res, err = h.EventStore.PostWithdraw(req)
		if err != nil {
			return utils.GetDefaultErrorMessage(c)
		}
		return c.JSON(http.StatusCreated, res)
	case types.DepositRequestTypeTransfer:
		res := &types.TransferResponse{}
		res, err = h.EventStore.PostTransfer(req)
		if err != nil {
			return utils.GetDefaultErrorMessage(c)
		}
		return c.JSON(http.StatusCreated, res)
	}

	return utils.GetDefaultErrorMessage(c)
}