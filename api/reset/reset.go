package reset

import "dt-tst/utils/mock_data"

type Handler struct{
	MockDataHandler *mock_data.MockDataHandler
}

// NewHandler create a new Reset Handler
func NewHandler(mockDataHandler *mock_data.MockDataHandler) *Handler {
	return &Handler{
		MockDataHandler: mockDataHandler,
	}
}

func (h *Handler) Reset() (res bool, err error){
	err = h.MockDataHandler.Reset()
	if err == nil {
		res = true
	}
	return
}