package api

type Usecase interface {
	FetchCountMessage() (int, error)
	UpdateCountMessage(msg int) error
}
