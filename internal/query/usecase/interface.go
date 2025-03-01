package usecase

type Provider interface {
	SelectQuery() (string, error)
	UpdateQuery(string) error
}
