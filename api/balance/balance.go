package balance

import "dt-tst/utils/mock_data"

type Handler struct{
	MockDataHandler *mock_data.MockDataHandler
}

// NewHandler create a new balance Handler
func NewHandler(mockDataHandler *mock_data.MockDataHandler) *Handler {
	return &Handler{
		MockDataHandler: mockDataHandler,
	}
}

func (h *Handler) GetBalance(accountId string) (balance float64, err error) {
	agent, err := h.MockDataHandler.GetAgentById(accountId)
	if err != nil {
		return
	}
	balance = agent.Balance
	return
}
