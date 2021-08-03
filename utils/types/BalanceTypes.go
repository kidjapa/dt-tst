package types

type (
	BalanceRequest struct {
		AccountId string `json:"account_id" form:"account_id" query:"account_id"`
	}
)
