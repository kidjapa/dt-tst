package interfaces

type ResetApiInterface interface {
	Reset() (bool, error)
}