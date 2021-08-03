package interfaces

type BalanceApiInterface interface {
	GetBalance(accountId string) (float64, error)
}