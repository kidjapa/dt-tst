package event

import (
	"dt-tst/utils/mock_data"
	"dt-tst/utils/types"
)

type Handler struct{
	MockDataHandler *mock_data.MockDataHandler
}

// NewHandler create a new Event Handler
func NewHandler(mockDataHandler *mock_data.MockDataHandler) *Handler {
	return &Handler{
		MockDataHandler: mockDataHandler,
	}
}

func (h *Handler) PostDeposit(depositRequest *types.EventRequest) (res *types.DepositResponse, err error) {
	res = &types.DepositResponse{}
	resAgent := &mock_data.Agent{}
	resAgent, err = h.MockDataHandler.DepositInAgent(depositRequest.Destination, depositRequest.Amount)
	if err != nil {
		return
	}
	res.Destination.Balance = resAgent.Balance
	res.Destination.Id = resAgent.AccountId
	return
}

func (h *Handler) PostWithdraw(depositRequest *types.EventRequest) (res *types.WithdrawResponse, err error) {
	res = &types.WithdrawResponse{}

	resAgent := &mock_data.Agent{}
	resAgent, err = h.MockDataHandler.WithdrawFromAgent(depositRequest.Origin, depositRequest.Amount)
	if err != nil {
		return
	}

	res.Origin.Balance = resAgent.Balance
	res.Origin.Id = resAgent.AccountId
	return
}

func (h *Handler) PostTransfer(depositRequest *types.EventRequest) (res *types.TransferResponse, err error) {
	res = &types.TransferResponse{}

	agent := &mock_data.Agent{}
	agent, err = h.MockDataHandler.GetAgentById(depositRequest.Origin)
	if err != nil {
		return
	}

	resAgent := &mock_data.Agent{}
	resAgent, err = h.MockDataHandler.TransferInAgent(agent, depositRequest.Destination, depositRequest.Amount)
	if err != nil {
		return
	}

	res.Origin.Balance = agent.Balance
	res.Origin.Id = agent.AccountId
	res.Destination.Balance = resAgent.Balance
	res.Destination.Id = resAgent.AccountId
	return
}