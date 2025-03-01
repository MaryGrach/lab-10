package usecase

func (u *Usecase) FetchQueryMessage() (string, error) {
	msg, err := u.p.SelectQuery()
	if err != nil {
		return "", err
	}

	return u.startMsg + msg + u.endMsg, nil
}

func (u *Usecase) SetQueryMessage(msg string) error {

	err := u.p.UpdateQuery(msg)
	if err != nil {
		return err
	}

	return nil
}
