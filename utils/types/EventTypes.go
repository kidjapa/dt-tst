package types

const (
	DepositRequestTypeDeposit  EventRequestType = "deposit"
	DepositRequestTypeWithDraw EventRequestType = "withdraw"
	DepositRequestTypeTransfer EventRequestType = "transfer"
)

type (
	EventRequestType string

	EventRequest struct {
		Type        EventRequestType `query:"type" form:"type" json:"type" validate:"required"`
		Destination string           `json:"destination"`
		Origin      string           `json:"origin"`
		Amount      float64          `json:"amount"`
	}

	EventCommonResponse struct {
		Id      string  `json:"id"`
		Balance float64 `json:"balance"`
	}

	DepositResponse struct {
		Destination EventCommonResponse `json:"destination,omitempty"`
	}

	WithdrawResponse struct {
		Origin EventCommonResponse `json:"origin,omitempty"`
	}

	TransferResponse struct {
		Origin      EventCommonResponse `json:"origin,omitempty"`
		Destination EventCommonResponse `json:"destination,omitempty"`
	}

	PostEventCommonResponse struct {
		Origin      EventCommonResponse `json:"origin,omitempty"`
		Destination EventCommonResponse `json:"destination,omitempty"`
	}
)
