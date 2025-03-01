package usecase

type Usecase struct {
	startMsg, endMsg string

	p Provider
}

func NewUsecase(p Provider, start_message, end_message string) *Usecase {
	return &Usecase{
		startMsg: start_message,
		endMsg:   end_message,
		p:        p,
	}
}
