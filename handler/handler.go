package handler

import "dt-tst/interfaces"

type Handler struct {
	BalanceStore interfaces.BalanceApiInterface
	EventStore   interfaces.EventApiInterface
	Reset        interfaces.ResetApiInterface
}

func NewHandler(es interfaces.EventApiInterface, bs interfaces.BalanceApiInterface, r interfaces.ResetApiInterface) *Handler {
	return &Handler{
		BalanceStore: bs,
		EventStore:   es,
		Reset: r,
	}
}
