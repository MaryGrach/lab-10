package api

type Usecase interface {
	FetchQueryMessage() (string, error)
	SetQueryMessage(msg string) error
}
