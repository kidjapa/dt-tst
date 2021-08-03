package interfaces

import "dt-tst/utils/types"

type EventApiInterface interface {
	PostDeposit(depositRequest *types.EventRequest) (*types.DepositResponse, error)
	PostWithdraw(depositRequest *types.EventRequest) (*types.WithdrawResponse, error)
	PostTransfer(depositRequest *types.EventRequest) (*types.TransferResponse, error)
}
