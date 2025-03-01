package usecase

type Provider interface {
	SelectCount() (int, error)
	UpdateCount(int) error
}
